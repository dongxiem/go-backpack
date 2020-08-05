package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ginEngine() *gin.Engine {
	router := gin.Default()
	// 输入http://localhost:8080/welcome
	router.GET("/welcome", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstName", "Guest")
		lastName := c.Query("lastName")

		c.String(http.StatusOK, "Hello %s %s", firstName, lastName)
	})
	// 输入http://localhost:8080/ping
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": gin.H{"data1": "data2"}})
	})

	return router
}

func main() {
	r := ginEngine()
	r.Run(":8080")
}
