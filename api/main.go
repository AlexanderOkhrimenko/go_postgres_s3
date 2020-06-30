package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_postgres_s3/api/modules"
)

func main() {

	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/url.insert", InsertUrl)

	}
	router.Run(":8080")

}

func InsertUrl(c *gin.Context) {

	id := modules.InsertDBurl("ya.ru")
	fmt.Println(id)
}
