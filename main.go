package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func serve(router *gin.Engine) (*gin.Engine, *http.Server) {
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

	srv := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	return router, srv
}

func main() {
	router := gin.Default()

	_, srv := serve(router)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
