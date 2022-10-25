package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func handlers() {
	// start each handler
	loadStatus()
	login()
	checkCookie()
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

					sessionToken := uuid.New().String()
					expiresAt := time.Now().Add(5 * time.Minute)

					_, err = db.Exec("INSERT INTO sessions(sessionname, username, expiration) VALUES ($1, $2, $3);", sessionToken, creds.Username, expiresAt)
					if err != nil {
						panic(err)
					}

					c.SetCookie("sessionToken", sessionToken, 300, "/", os.Getenv("DOMAIN"), false, true)

					// return valid status if passwords match
					c.JSON(http.StatusOK, gin.H{
						"status": "Valid Credentials",
					})

				}
			}
		}
	})
}

func checkCookie() {
	router.GET("/backend/checkCookie", func(c *gin.Context) {
		cookie, err := c.Cookie("sessionToken")
		if err != nil {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		var expiration time.Time
		query := fmt.Sprintf("SELECT expiration FROM sessions WHERE sessionname = '%s';", cookie)
		row := db.QueryRow(query)
		err = row.Scan(&expiration)
		if err != nil {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		if expiration.Before(time.Now()) {
			query = fmt.Sprintf("DELETE FROM sessions WHERE sessioname = '%s';", cookie)
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		sessionToken := uuid.New().String()
		var username string
		expiresAt := time.Now().Add(5 * time.Minute)

		query = fmt.Sprintf("SELECT username FROM sessions WHERE sessionname = '%s';", cookie)
		row = db.QueryRow(query)
		err = row.Scan(&username)
		if err != nil {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		_, err = db.Exec("INSERT INTO sessions(sessionname, username, expiration) VALUES ($1, $2, $3);", sessionToken, username, expiresAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			panic(err)
		}

		_, err = db.Exec("DELETE FROM sessions WHERE sessionname = $1;", cookie)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			panic(err)
		}

		c.SetCookie("sessionToken", sessionToken, 300, "/", os.Getenv("DOMAIN"), false, true)

		c.JSON(http.StatusOK, nil)
		return
	})
}
