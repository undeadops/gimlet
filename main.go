// main.go

package main

import (
	"fmt"
	"log"
	//"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/undeadops/gimlet/healthz"
)

var version = "2.0.0"

func index(c *gin.Context) {
	hostname, _ := os.Hostname()
	content := gin.H{"version": version, "host": hostname}
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
	router.GET("/healthz", healthz.Status)
	fmt.Printf("HTTP service listening on %s\n", httpAddr)
	router.Run(":" + httpAddr)
}
