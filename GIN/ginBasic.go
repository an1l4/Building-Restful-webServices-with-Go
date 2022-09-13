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

curl -X POST \http://localhost:8000/v1/stations \-H 'cache-control: no-cache' \-H 'content-type: application/json' \-d '{"name":"Brooklyn", "opening_time":"8:12:00","closing_time":"18:23:00"}'
