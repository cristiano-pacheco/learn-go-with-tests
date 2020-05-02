package main

import "github.com/gin-gonic/gin"

func Home(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Oh Yeah",
	})
}

func main() {
	router := gin.Default()

	router.GET("/", Home)

	router.Run()
}
