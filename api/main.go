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
	// https://levelup.gitconnected.com/dockerized-crud-restful-api-with-go-gorm-jwt-postgresql-mysql-and-testing-61d731430bd8
	// https://levelup.gitconnected.com/crud-restful-api-with-go-gorm-jwt-postgres-mysql-and-testing-460a85ab7121
	// https://dev.to/stevensunflash/real-world-app-with-golang-gin-and-react-hooks-44ph
	// https://github.com/victorsteven/Go-JWT-Postgres-Mysql-Restful-API

}

func InsertUrl(c *gin.Context) {

	id := modules.InsertDBurl("ya.ru")
	c.String(200, string(id))

}
