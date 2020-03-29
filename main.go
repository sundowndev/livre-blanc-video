package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func serve(router *gin.Engine) *gin.Engine {
	router.Use(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"success": false,
			"error":   "Resource not found",
		})
	})

	return router
}

func main() {
	router := gin.Default()

	serve(router)

	srv := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
