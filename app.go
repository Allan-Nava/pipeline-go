/*
#	Author		: 	Allan Nava
#	Modified	:	Allan Nava
#	Date		:	30/04/2019
#	Updated 	: 	30/04/2019
*/
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
