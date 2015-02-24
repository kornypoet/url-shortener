package main

import (
  "crypto/md5"
  "encoding/hex"
  "fmt"
  "net/url"
  "github.com/gin-gonic/gin"
  // "gopkg.in/mgo.v2"
  // "gopkg.in/mgo.v2/bson"
)

type JSONBody struct {
  Url string `json:"url" binding:"required"`
}

func main() {
  router := gin.Default()
  router.Use(gin.Logger())

  router.GET("/" , func(c *gin.Context) {
    c.String(200, "Hello World")
  })

  router.POST("/shorten", func(c *gin.Context) {
    var json JSONBody

    c.Bind(&json)

    url, err := url.Parse(json.Url)
    if err != nil || url.Host == "" {
      fmt.Println(err)
      c.String(400, "Malformed Url")
    } else {

      hasher := md5.New()
      hasher.Write([]byte(json.Url))
      id := hex.EncodeToString(hasher.Sum(nil))

      var msg struct {
        Path  string `json:"path"`
        Count int    `json:"count"`
        Host  string `json:"host"`
      }

      msg.Path  = id
      msg.Host  = url.Host
      msg.Count = 1

      c.JSON(200, msg)
    }
  })

  router.Run(":8080")
}
