package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"DataSyncServer/database"
	"DataSyncServer/entities"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	r.GET("/api/messages", get)
	r.POST("/api/messages", post)
	r.PUT("/api/messages/:uuid", put)
	r.DELETE("/api/messages/:uuid", delete)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func get(c *gin.Context) {
	var requestDatetime entities.GetRequestBody
	c.BindJSON(&requestDatetime)
	fmt.Println(requestDatetime.UnixTimestamp)
	tm := time.Unix(int64(requestDatetime.UnixTimestamp), 0).UTC()
	fmt.Println(tm)
}

func post(c *gin.Context) {
	var postRequestBody entities.PostRequestBody
	c.BindJSON(&postRequestBody)
	db := database.Connect()
	defer db.Close()
	insForm, err := db.Prepare("INSERT INTO data_record(uuid,author,message,likes,created_at) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	tm := time.Now().UTC()
	_, err = insForm.Exec(
		postRequestBody.UUID,
		postRequestBody.Author,
		postRequestBody.Message,
		postRequestBody.Likes,
		tm,
	)
	if err != nil {
		e, _ := err.(*mysql.MySQLError)
		// if uuid already exists
		if e.Number == 1062 {
			c.JSON(409, nil)
			return
		}
	}
	// on success
	c.JSON(201, nil)
}

func put(c *gin.Context) {
	uuid := c.Param("uuid")
	var putRequestBody entities.PutRequestBody
	c.BindJSON(&putRequestBody)
	db := database.Connect()
	defer db.Close()
	row := db.QueryRow("SELECT author,message,likes FROM data_record WHERE uuid=?", uuid)
	var author, message string
	var likes int
	err := row.Scan(&author, &message, &likes)
	if err != nil {
		if err == sql.ErrNoRows {
			// if uuid is not found
			c.JSON(404, nil)
			return
		}
	}
	var query []string
	var values []interface{}
	tm := time.Now().UTC()
	change := false
	if author != putRequestBody.Author {
		query = append(query, "author = ?")
		query = append(query, "author_updated_at = ?")
		values = append(values, putRequestBody.Author)
		values = append(values, tm)
		change = true
	}
	if message != putRequestBody.Message {
		query = append(query, "message = ?")
		query = append(query, "message_updated_at = ?")
		values = append(values, putRequestBody.Message)
		values = append(values, tm)
		change = true

	}
	if likes != putRequestBody.Likes {
		query = append(query, "likes = ?")
		query = append(query, "likes_updated_at = ?")
		values = append(values, putRequestBody.Likes)
		values = append(values, tm)
		change = true
	}
	if change {
		update_sql := fmt.Sprintf(`UPDATE data_record SET %v, updated=1 WHERE uuid="%v"`, strings.Join(query, ", "), uuid)
		insForm, err := db.Prepare(update_sql)
		if err != nil {
			panic(err.Error())
		}
		_, err = insForm.Exec(values...)
		if err != nil {
			panic(err.Error())
		}
	}
	// on success
	c.JSON(204, nil)
}

func delete(c *gin.Context) {
	uuid := c.Param("uuid")
	db := database.Connect()
	defer db.Close()
	row := db.QueryRow("SELECT author,message,likes FROM data_record WHERE uuid=?", uuid)
	var author, message string
	var likes int
	err := row.Scan(&author, &message, &likes)
	if err != nil {
		if err == sql.ErrNoRows {
			// if uuid is not found
			c.JSON(404, nil)
			return
		}
	}
	// on success
	c.JSON(204, nil)
	return
	// if uuid is not found
	c.JSON(404, nil)
}
