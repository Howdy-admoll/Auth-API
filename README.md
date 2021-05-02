# Auth-API
Creating an Authentication API with Golang

<img align="right" width="159px" src="https://raw.githubusercontent.com/assertgo/icon/master/assertgo_512.png">

[![GoDoc](https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg)](https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc)

# Installation

You need to install Go on your PC or MacOS and set your Go workspace first

- The first need [Go](https://golang.org/) installed (**Latest Stable version is awesome**)

# Getting Started

- To get started, Create a new project then, Create a new file called `main.go` and enter the following starter code:

```go
package main
import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)
```

# POST and GET methods

- To be able to handle GET and POST requests, we create a `GetMethod` and `PostMethod`:

```go
func PostMethod(c *gin.Context) {
  fmt.Println("\napi.go 'PostMethod' called")
  message := "PostMethod called"
  c.JSON(http.StatusOK, message)
}
func GetMethod(c *gin.Context) {
  fmt.Println("\napi.go 'GetMethod' called")
  message := "GetMethod called"
  c.JSON(http.StatusOK, message)
}
```
# Main function

- We need to create a `Gin` router for handling all the requests:

```go
func main() {
  router := gin.Default()
}
```

- Following this up, we add in the `GetMethod` and `PostMethod`:

```go
func main() {
  router := gin.Default()
router.POST("/", PostMethod)
  router.GET("/", GetMethod)
}
```
