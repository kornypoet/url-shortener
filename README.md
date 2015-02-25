# URL Shortener

## Requirements

Go dependencies:

```
go get github.com/gin-gonic/gin
go get gopkg.in/mgo.v2
```

A running MongoDB instance:

```
mongod
```

## Installation

```
git clone git@github.com:kornypoet/url-shortener.git
cd url-shortener
go install url-shortener.go
```

## Usage

```
url-shortener
```

To check the status of the server:

```
curl -X GET 'http://localhost:8080/status?'
```

To create an entry:

```
curl -X POST -H 'Content-Type: application/json' 'http://localhost:8080/shorten?' -d '{"url":"https://www.youtube.com/watch?v=dQw4w9WgXcQ"}'
# {"url":"https://www.youtube.com/watch?v=dQw4w9WgXcQ","path":"75170fc230cd88f32e475ff4087f81d9","count":0,"host":"www.youtube.com"}
```

To access the shortened url:

```
curl -X GET -I 'http://localhost:8080/shorten/75170fc230cd88f32e475ff4087f81d9'
# HTTP/1.1 301 Moved Permanently
# Location: https://www.youtube.com/watch?v=dQw4w9WgXcQ
# Date: Wed, 25 Feb 2015 18:31:32 GMT
# Content-Length: 0
# Content-Type: text/plain; charset=utf-8
```

And to retrieve data about the shortened url:

```
curl -X GET 'http://localhost:8080/info/75170fc230cd88f32e475ff4087f81d9'
# {"url":"https://www.youtube.com/watch?v=dQw4w9WgXcQ","path":"75170fc230cd88f32e475ff4087f81d9","count":1,"host":"www.youtube.com"}
```
