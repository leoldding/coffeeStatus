package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{":3000"})

	initDB()

	router.Run(":8080")

	return
}
