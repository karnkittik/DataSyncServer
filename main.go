package main

import (
	"fmt"
	"log"

	"DataSyncServer/database"
	"DataSyncServer/entities"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")
	db := database.Connect()
	rows, err := db.Query("SELECT id, username FROM my_table")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		var users []entities.User
		for rows.Next() {
			var user entities.User
			err := rows.Scan(&user.Id, &user.Username)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, user)
		}
		c.JSON(200, users)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// run go run main.go and visit 0.0.0.0:8080/ping (for windows "localhost:8080/ping") on browser
}
