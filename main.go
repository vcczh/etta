package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/health_check", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	router.Run(":8080")
}
