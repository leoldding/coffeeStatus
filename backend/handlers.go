package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
)

func handlers() {
	// start each handler
	loadStatus()
	login()
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loadStatus() {
	router.GET("/backend/loadStatus", func(c *gin.Context) {
		// retrieve status from database
		var status string
		query := fmt.Sprintf("SELECT status FROM admins WHERE adminname = '%s';", os.Getenv("ADMINNAME"))
		row := db.QueryRow(query)
		err = row.Scan(&status)
		if err != nil {
			panic(err)
		}

		// return JSON data to front with current status
		c.JSON(http.StatusOK, gin.H{
			"status": status,
		})
	})
}

func login() {
	router.POST("/backend/login", func(c *gin.Context) {
		var creds Credentials
		if c.BindJSON(&creds) == nil {
			var count int
			// check if username exists in database
			query := fmt.Sprintf("SELECT COUNT(*) FROM admins WHERE adminname = '%s';", creds.Username)
			row := db.QueryRow(query)
			err = row.Scan(&count)
			if err != nil {
				panic(err)
			}
			if count == 0 {
				// return unauthorized status if username doesn't exist
				c.JSON(http.StatusUnauthorized, gin.H{
					"status": "Invalid Username",
				})
			} else {
				var password []byte
				// retrieve password from database with valid username
				query = fmt.Sprintf("SELECT password FROM admins WHERE adminname = '%s';", creds.Username)
				row = db.QueryRow(query)
				err = row.Scan(&password)
				if err != nil {
					panic(err)
				}

				// compare inputted password with stored password hash
				err = bcrypt.CompareHashAndPassword(password, []byte(creds.Password))
				if err != nil {
					// return unauthorized status if passwords don't match
					c.JSON(http.StatusUnauthorized, gin.H{
						"status": "Invalid Password",
					})
				} else {
					// return valid status if passwords match
					c.JSON(http.StatusOK, gin.H{
						"status": "Valid Credentials",
					})
				}
			}
		}
	})
}
