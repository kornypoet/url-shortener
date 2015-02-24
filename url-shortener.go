package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

  router.GET("/" , func(c *gin.Context) {
    c.String(200, "Hello World")
  })

	router.GET("/shorten/:path", func(c *gin.Context) {
    c.String(200, "Shortening path " + c.Params.ByName("path"))
  })

  router.Run(":8080")
}
