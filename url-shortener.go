package main

import(
  "crypto/md5"
  "encoding/hex"
  "fmt"
  "net/url"
  "github.com/gin-gonic/gin"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type JSONBody struct {
  Url   string `json:"url" binding:"required"`
  Path  string `json:"path"`
  Count int    `json:"count"`
  Host  string `json:"host"`
}

type UrlDoc struct {
  Id    string `bson:"_id"`
  Url   string
  Host  string
  Count int
}

func createId(s string) string {
  hasher := md5.New()
  hasher.Write([]byte(s))
  return hex.EncodeToString(hasher.Sum(nil))
}

func updateRecord(id string, coll *mgo.Collection) (result UrlDoc, err error) {
  result  = UrlDoc{}
  change := mgo.Change{
    Update:    bson.M{"$inc": bson.M{"count": 1}},
    ReturnNew: true,
  }
  _, err = coll.Find(bson.M{"_id": id}).Apply(change, &result)
  return result, err
}

func findOrCreate(doc UrlDoc, coll *mgo.Collection) (result UrlDoc, err error) {
  result = UrlDoc{}
  findErr := coll.Find(bson.M{"_id": doc.Id}).One(&result)
  if findErr != nil {
    fmt.Println("Record not found: creating new entry")
    err = coll.Insert(&doc)
    result = doc
  }
  return result, err
}

func main() {
  mongo, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  defer mongo.Close()
  mongo.SetMode(mgo.Monotonic, true)
  coll := mongo.DB("test").C("urls")

  router := gin.Default()
  router.Use(gin.Logger())

  router.GET("/status" , func(c *gin.Context) {
    c.String(200, "OK")
  })

  router.GET("/shorten/:id", func(c *gin.Context) {
    result, err := updateRecord(c.Params.ByName("id"), coll)
    if err != nil {
      c.String(404, "Not Found")
    } else {
      c.Redirect(301, result.Url)
    }
  })

  router.POST("/shorten", func(c *gin.Context) {
    var json JSONBody
    c.Bind(&json)
    url, err := url.Parse(json.Url)
    if err != nil || url.Host == "" {
      fmt.Println(err)
      c.String(400, "Malformed Url")
    } else {
      id := createId(json.Url)
      doc := UrlDoc{id, json.Url, url.Host, 1}
      result, err := findOrCreate(doc, coll)
      if err != nil {
        fmt.Println(err)
      }
      msg := JSONBody{result.Url, result.Id, result.Count, result.Host}
      c.JSON(200, msg)
    }
  })

  router.Run(":8080")
}
