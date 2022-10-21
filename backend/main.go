package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// startup router
	router = gin.Default()
	router.SetTrustedProxies([]string{":3000"})

	// initialize databases
	initDB()

	// call all handlers
	handlers()
	
	router.Run(":8080")
}
