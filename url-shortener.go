package main

import (
	"github.com/go-martini/martini"
)

func main() {
  m := martini.Classic()

  m.Get("/" , func() string {
    return "Hello World"
  })

	m.Get("/shorten/:path", func(params martini.Params) string {
    return "Shortening path " + params["path"]
  })

  m.Run()
}
