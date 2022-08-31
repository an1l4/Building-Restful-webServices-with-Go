package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//	log.Println("check1")
	r.GET("/pingTime", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})
	//log.Printf("Server stopped, err: %v", r.Run(":8080"))

	r.Run(":8080") // Default listen and serve on 0.0.0.0:8080
}
