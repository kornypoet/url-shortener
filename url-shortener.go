package main

import (
  "crypto/md5"
  "encoding/hex"
  "fmt"
  "net/url"
  "github.com/gin-gonic/gin"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type JSONBody struct {
  Url string `json:"url" binding:"required"`
}

type Payload struct {
  Id    string `bson:"_id"`
  Url   string
  Count int
}


func main() {

  mongo, err := mgo.Dial("localhost")
  if err != nil {
    fmt.Println(err)
  }

  defer mongo.Close()
  mongo.SetMode(mgo.Monotonic, true)

  coll := mongo.DB("test").C("urls")

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

      result := Payload{}
      err = coll.Find(bson.M{"_id": id}).One(&result)
      if err != nil {
        fmt.Println(err)

        err = coll.Insert(&Payload{id, json.Url, 1})
        if err != nil {
          fmt.Println(err)
        }
      }

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
