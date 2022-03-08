package main

import (
	"fmt"
	"log"
	"time"

	"DataSyncServer/database"
	"DataSyncServer/entities"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()

	r := gin.Default()
	r.GET("/api/messages", get)
	r.POST("/api/messages", post)
	r.PUT("/api/messages/:uuid", put)
	r.DELETE("/api/messages/:uuid", delete)

	r.GET("/users", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, username FROM my_table")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		var users []entities.DataRecord
		for rows.Next() {
			var user entities.DataRecord
			err := rows.Scan(&user.UUID, &user.Author)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, user)
		}
		c.JSON(200, users)
	})
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
	// on success
	c.JSON(201, nil)
	return
	// if uuid already exists
	c.JSON(409, nil)
}

func put(c *gin.Context) {
	uuid := c.Param("uuid")
	fmt.Println(uuid)
	var putRequestBody entities.PutRequestBody
	c.BindJSON(&putRequestBody)
	// on success
	c.JSON(204, nil)
	return
	// if uuid is not found
	c.JSON(404, nil)
}

func delete(c *gin.Context) {
	uuid := c.Param("uuid")
	fmt.Println(uuid)
	// on success
	c.JSON(204, nil)
	return
	// if uuid is not found
	c.JSON(404, nil)
}
