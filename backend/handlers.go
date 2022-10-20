package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func handlers() {
	backend := router.Group("/backend")
	{
		backend.GET("/loadstatus", func(c *gin.Context) {
			var status string
			query := fmt.Sprintf("SELECT status FROM admins WHERE adminname = '%s';", os.Getenv("ADMINNAME"))
			fmt.Print("status select query: ")
			fmt.Println(query)
			row := db.QueryRow(query)
			err = row.Scan(&status)
			if err != nil {
				log.Fatal(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"status": status,
			})
		})
	}
}
