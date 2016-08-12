// main.go

package main

import (
	"fmt"
	"log"
	//"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var version = "1.0.0"

func index(c *gin.Context) {
	content := gin.H{"version": version}
	c.JSON(200, content)
}

func main() {
	log.Println("Starting gimlet... ")
	gin.SetMode(gin.ReleaseMode)
	httpAddr := os.Getenv("PORT")
	if httpAddr == "" {
		httpAddr = "8080"
	}
	router := gin.Default()

	router.GET("/", index)
	fmt.Printf("HTTP service listening on %s\n", httpAddr)
	router.Run(":" + httpAddr)
}
