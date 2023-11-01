# RoutiX
🚀 Routing X is a Golang library for routing REST APIs 🚀

# Features
- Build for [Gin gonic](https://github.com/gin-gonic/gin)
- Auto grouping router on instance gin engine
- Auto apply middlewares for global
- Just return any data type from the handler function without using additional functions
- Return HTTP exception filter as JSON status code and message
- Guard
- Interceptor

  
# Installation
```
go get github.com/l1ttps/routix

go get github.com/gin-gonic/gin
```

# Example of usage
1. Create a file app_controler.go
```go
package controller

import (
 "github.com/gin-gonic/gin"
 "github.com/l1ttps/routix"
)

var (
  Controller   = routix.Controller
  Get          = routix.Get
)

func AppController() *gin.Engine {
  return Controller("/",
    Get("/",
      func(c *gin.Context) interface{} {
        return "hello world"
      },
    ),
  )
}
```

2. Create a file main.go
```go
package main

import (
 "github.com/l1ttps/routix"
 "<your_pakage>/controller"
)

var (
  CreateServer = routix.CreateServer
)

func main() {

 app := CreateServer(routix.ServerConfig{
    Controllers: []routix.ControllerType{
      controller.AppController,
    },
     DebugLogger: true,
    })

  app.Run(":3000")
}

```
3. Run go run main.go and open http://localhost:3000
