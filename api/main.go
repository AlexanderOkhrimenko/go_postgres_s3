package main

import (

	"github.com/gin-gonic/gin"
	"go_postgres_s3/api/modules"
)

func main() {

	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	router.MaxMultipartMemory = 8 << 30 // 8 MiB

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/url.insert", InsertUrl) // Upload MultipartForm

	}

	router.Run(":8080")

}

func InsertUrl(c *gin.Context) {

	id := modules.InsertDBurl("ya.ru")
		c.String(200, string(id))

}
