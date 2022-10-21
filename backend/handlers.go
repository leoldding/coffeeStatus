package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func handlers() {
	// start each handler
	loadstatus()
}

func loadstatus() {
	router.GET("/backend/loadstatus", func(c *gin.Context) {
		// retrieve status from database
		var status string
		query := fmt.Sprintf("SELECT status FROM admins WHERE adminname = '%s';", os.Getenv("ADMINNAME"))
		row := db.QueryRow(query)
		err = row.Scan(&status)
		if err != nil {
			log.Fatal(err)
		}

		// return JSON data to front with current status
		c.JSON(http.StatusOK, gin.H{
			"status": status,
		})
	})
}
