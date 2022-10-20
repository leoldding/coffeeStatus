package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.SetTrustedProxies([]string{":3000"})

	initDB()

	handlers()

	router.Run(":8080")
}
