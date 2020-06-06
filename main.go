package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) *gin.Engine {
	router.StaticFile("/", "./public")

	router.GET("/api/dataset", func(c *gin.Context) {
		csv, err := ioutil.ReadFile("./data/responses.csv")
		if err != nil {
			log.Fatal(err)
		}

		c.Header("Content-Type", "text/plain; charset=utf-8")
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte(csv))
		c.Abort()
	})

	router.Use(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Resource not found",
		})
	})

	return router
}

func main() {
	router := gin.Default()

	registerRoutes(router)

	router.Run(":5000")
}
