package main

import (
	"fmt"
	customer "fullstack/controller"
	"fullstack/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	db.Connect()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	r.GET("/get-customers", customer.GetCustomers)
	fmt.Println("http://localhost:8080")
	r.Run(":8080")
}
