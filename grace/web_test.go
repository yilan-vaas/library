package grace

import (
	"log"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGraceWeb(t *testing.T) {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
		return
	})
	server := NewServer("", r)
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
