package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	// assign port for running server
	port := os.Getenv("PORT")

	// if port is blank then assign it default server
	if port == "" {
		port = "8000"
	}

	// make router instance
	router := gin.New()

	// assign logger to router instance
	router.Use(gin.Logger())

}
