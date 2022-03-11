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

func select_create(db *sql.DB, tm time.Time) []entities.ResponseData {
	selDB, err := db.Query("SELECT uuid,author,message,likes FROM data_record WHERE created_at > ? AND deleted=0", tm)
	if err != nil {
		panic(err.Error())
	}
	records := []entities.ResponseData{}
	for selDB.Next() {
		var rec entities.ResponseData
		err = selDB.Scan(&rec.UUID, &rec.Author, &rec.Message, &rec.Likes)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, rec)
	}
	return records
}
func select_update_author(db *sql.DB, tm time.Time) []entities.ResponseData {
	selDB, err := db.Query("SELECT uuid,author FROM data_record WHERE author_updated_at > ? AND (message_updated_at <= ? OR message_updated_at IS NULL) AND (likes_updated_at <= ? OR likes_updated_at IS NULL) AND updated=1 AND deleted=0", tm, tm, tm)
	if err != nil {
		panic(err.Error())
	}
	records := []entities.ResponseData{}
	for selDB.Next() {
		var rec entities.ResponseData
		err = selDB.Scan(&rec.UUID, &rec.Author)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, rec)
	}
	return records
}
func select_update_message(db *sql.DB, tm time.Time) []entities.ResponseData {
	selDB, err := db.Query("SELECT uuid,message FROM data_record WHERE message_updated_at > ? AND (author_updated_at <= ? OR author_updated_at IS NULL) AND (likes_updated_at <= ? OR likes_updated_at IS NULL) AND updated=1 AND deleted=0", tm, tm, tm)
	if err != nil {
		panic(err.Error())
	}
	records := []entities.ResponseData{}
	for selDB.Next() {
		var rec entities.ResponseData
		err = selDB.Scan(&rec.UUID, &rec.Message)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, rec)
	}
	return records
}
func select_update_likes(db *sql.DB, tm time.Time) []entities.ResponseData {
	selDB, err := db.Query("SELECT uuid,likes FROM data_record WHERE likes_updated_at > ? AND (author_updated_at <= ? OR author_updated_at IS NULL) AND (message_updated_at <= ? OR message_updated_at IS NULL) AND updated=1 AND deleted=0", tm, tm, tm)
	if err != nil {
		panic(err.Error())
	}
	records := []entities.ResponseData{}
	for selDB.Next() {
		var rec entities.ResponseData
		err = selDB.Scan(&rec.UUID, &rec.Likes)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, rec)
	}
	return records
}
func select_update_author_message(db *sql.DB, tm time.Time) []entities.ResponseData {
	selDB, err := db.Query("SELECT uuid,author,message FROM data_record WHERE author_updated_at > ? AND message_updated_at > ? AND (likes_updated_at <= ? OR likes_updated_at IS NULL) AND updated=1 AND deleted=0", tm, tm, tm)
	if err != nil {
		panic(err.Error())
	}
	records := []entities.ResponseData{}
	for selDB.Next() {
		var rec entities.ResponseData
		err = selDB.Scan(&rec.UUID, &rec.Author, &rec.Message)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, rec)
	}
	return records
}
func select_update_author_likes(db *sql.DB, tm time.Time) []entities.ResponseData {
	selDB, err := db.Query("SELECT uuid,author,likes FROM data_record WHERE author_updated_at > ? AND likes_updated_at > ? AND (message_updated_at <= ? OR message_updated_at IS NULL) AND updated=1 AND deleted=0", tm, tm, tm)
	if err != nil {
		panic(err.Error())
	}
	records := []entities.ResponseData{}
	for selDB.Next() {
		var rec entities.ResponseData
		err = selDB.Scan(&rec.UUID, &rec.Author, &rec.Likes)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, rec)
	}
	return records
}
func select_update_message_likes(db *sql.DB, tm time.Time) []entities.ResponseData {
	selDB, err := db.Query("SELECT uuid,message,likes FROM data_record WHERE message_updated_at > ? AND likes_updated_at > ? AND (author_updated_at <= ? OR author_updated_at IS NULL) AND updated=1 AND deleted=0", tm, tm, tm)
	if err != nil {
		panic(err.Error())
	}
	records := []entities.ResponseData{}
	for selDB.Next() {
		var rec entities.ResponseData
		err = selDB.Scan(&rec.UUID, &rec.Message, &rec.Likes)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, rec)
	}
	return records
}
func select_update_auther_message_likes(db *sql.DB, tm time.Time) []entities.ResponseData {
	selDB, err := db.Query("SELECT uuid,author,message,likes FROM data_record WHERE author_updated_at > ? AND message_updated_at > ? AND likes_updated_at > ? AND updated=1 AND deleted=0", tm, tm, tm)
	if err != nil {
		panic(err.Error())
	}
	records := []entities.ResponseData{}
	for selDB.Next() {
		var rec entities.ResponseData
		err = selDB.Scan(&rec.UUID, &rec.Author, &rec.Message, &rec.Likes)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, rec)
	}
	return records
}

func select_delete(db *sql.DB, tm time.Time) []string {
	selDB, err := db.Query("SELECT uuid FROM data_record WHERE created_at > ? AND deleted=1", tm)
	if err != nil {
		panic(err.Error())
	}
	var records []string
	for selDB.Next() {
		var uuid string
		err = selDB.Scan(&uuid)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, uuid)
	}
	return records
}

func get(c *gin.Context) {
	var requestDatetime entities.GetRequestBody
	c.BindJSON(&requestDatetime)
	db := database.Connect()
	defer db.Close()
	fmt.Println(requestDatetime.UnixTimestamp)
	tm := time.Unix(int64(requestDatetime.UnixTimestamp), 0).UTC()
	fmt.Println(tm)
	create_list := select_create(db, tm)
	delete_list := select_delete(db, tm)
	update := []entities.ResponseData{}
	// go routine
	update_author_list := select_update_author(db, tm)
	update_message_list := select_update_message(db, tm)
	update_likes_list := select_update_likes(db, tm)
	update_author_message_list := select_update_author_message(db, tm)
	update_author_likes_list := select_update_author_likes(db, tm)
	update_message_likes_list := select_update_message_likes(db, tm)
	update_auther_message_likes_list := select_update_auther_message_likes(db, tm)
	// combine data
	update = append(update, update_author_list...)
	update = append(update, update_message_list...)
	update = append(update, update_likes_list...)
	update = append(update, update_author_message_list...)
	update = append(update, update_author_likes_list...)
	update = append(update, update_message_likes_list...)
	update = append(update, update_auther_message_likes_list...)
	m := map[string]interface{}{
		"create": create_list,
		"delete": delete_list,
		"update": update,
	}
	c.JSON(200, m)
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
		upForm, err := db.Prepare(update_sql)
		if err != nil {
			panic(err.Error())
		}
		_, err = upForm.Exec(values...)
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
	delete_sql := "UPDATE data_record SET deleted=1 WHERE uuid=?"
	delForm, err := db.Prepare(delete_sql)
	if err != nil {
		panic(err.Error())
	}
	_, err = delForm.Exec(uuid)
	if err != nil {
		panic(err.Error())
	}
	// on success
	c.JSON(204, nil)
}
