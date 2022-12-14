package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func handlers() {
	// start each handler
	loadStatus()
	updateStatus()
	login()
	logout()
	checkCookie()
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Update struct {
	Status    string `json:"status"`
	Substatus string `json:"substatus"`
}

func loadStatus() {
	router.GET("/backend/loadStatus", func(c *gin.Context) {
		// retrieve status from database
		var status string
		row := db.QueryRow("SELECT status FROM status;")
		err = row.Scan(&status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			panic(err)
			return
		}

		// retrieve substatus from database
		var substatus string
		row = db.QueryRow("SELECT substatus FROM status;")
		err = row.Scan(&substatus)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			panic(err)
			return
		}

		// return JSON data to front with current status
		c.JSON(http.StatusOK, gin.H{
			"status":    status,
			"substatus": substatus,
		})

		return
	})
}

func updateStatus() {
	router.POST("/backend/updateStatus", func(c *gin.Context) {
		// retrieve JSON information
		var update Update
		if c.BindJSON(&update) != nil {
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		// change values in database
		_, err = db.Exec("UPDATE status SET status = $1, substatus = $2;", update.Status, update.Substatus)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, nil)
		return
	})
}

func login() {
	router.POST("/backend/login", func(c *gin.Context) {
		// retrieve JSON information
		var creds Credentials
		if c.BindJSON(&creds) != nil {
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		var count int
		// check if username exists in database
		row := db.QueryRow("SELECT COUNT(*) FROM admins WHERE adminname = $1;", creds.Username)
		err = row.Scan(&count)
		if err != nil {
			panic(err)
		}
		if count == 0 {
			// return unauthorized status if username doesn't exist
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Invalid Username",
			})
			return
		}

		var password []byte
		// retrieve password from database with valid username
		row = db.QueryRow("SELECT password FROM admins WHERE adminname = $1;", creds.Username)
		err = row.Scan(&password)
		if err != nil {
			panic(err)
		}

		// compare inputted password with stored password hash
		err = bcrypt.CompareHashAndPassword(password, []byte(creds.Password))
		if err != nil {
			// return unauthorized status if passwords don't match
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Incorrect Password",
			})
			return
		}

		// set values for cookie
		sessionToken := uuid.New().String()
		expiresAt := time.Now().Add(5 * time.Minute)

		_, err = db.Exec("INSERT INTO sessions(sessionname, username, expiration) VALUES ($1, $2, $3);", sessionToken, creds.Username, expiresAt)
		if err != nil {
			panic(err)
		}

		// set new cookie
		c.SetCookie("sessionToken", sessionToken, 300, "/", "", false, true)

		// return valid status if passwords match
		c.JSON(http.StatusOK, gin.H{
			"status": "Valid Credentials",
		})
		return
	})
}

func logout() {
	router.GET("/backend/logout", func(c *gin.Context) {
		// check if cookie exists
		cookie, err := c.Cookie("sessionToken")
		if err != nil {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		// delete cookie from sessions database
		_, err = db.Exec("DELETE FROM sessions WHERE sessionname = $1;", cookie)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		// unset cookie within browser
		c.SetCookie("sessionToken", "", -1, "/", "", false, true)
		return
	})
}

func checkCookie() {
	router.GET("/backend/checkCookie", func(c *gin.Context) {
		// retrieve cookie
		cookie, err := c.Cookie("sessionToken")
		// check if cookie exists
		if err != nil {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		// check expiration time
		var expiration time.Time
		row := db.QueryRow("SELECT expiration FROM sessions WHERE sessionname = $1;", cookie)
		err = row.Scan(&expiration)
		if err != nil {
			// return unauthorized status if cookie doesn't exist in database
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		if expiration.Before(time.Now()) {
			// delete expired cookie and return unauthorized status
			db.Exec("DELETE FROM sessions WHERE sessioname = $1;", cookie)
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		// set new cookie values
		sessionToken := uuid.New().String()
		var username string
		expiresAt := time.Now().Add(5 * time.Minute)

		// put cookie into database
		row = db.QueryRow("SELECT username FROM sessions WHERE sessionname = $1;", cookie)
		err = row.Scan(&username)
		if err != nil {
			// return unauthorized status if cookie doesn't exist in database
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		// insert new cookie into database
		_, err = db.Exec("INSERT INTO sessions(sessionname, username, expiration) VALUES ($1, $2, $3);", sessionToken, username, expiresAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			panic(err)
		}

		// delete old cookie from database
		_, err = db.Exec("DELETE FROM sessions WHERE sessionname = $1;", cookie)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			panic(err)
		}

		// set new cookie in browser (overrides old cookie)
		c.SetCookie("sessionToken", sessionToken, 300, "/", "", false, true)

		c.JSON(http.StatusOK, nil)
		return
	})
}
