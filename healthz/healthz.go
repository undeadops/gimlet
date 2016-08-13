package healthz

import (
	"github.com/gin-gonic/gin"
)

func Status(c *gin.Context) {
	//Run Status checks
	content := gin.H{"status": "OK"}
	c.JSON(200, content)
}
