package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/pingTime", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})
	r.Run(":8080")
}
