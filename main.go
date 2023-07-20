package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", HealthCheckHandler)
	r.Run(":8080")
}

func HealthCheckHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}
