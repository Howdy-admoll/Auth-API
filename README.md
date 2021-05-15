# Auth-API
Creating an Authentication API with Golang

<img align="right" width="159px" src="https://raw.githubusercontent.com/assertgo/icon/master/assertgo_512.png">

[![GoDoc](https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg)](https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc)

# Installation (ref: Thanks to ==> Dijkhuizen)

You need to install Go on your PC or MacOS and set your Go workspace first

- The first need [Go](https://golang.org/) installed (**Latest {stable} version is awesome**)

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
- we add the port for the server to listen to and run the server: Just top the below code in `Gin`

```go
func main() {
  router := gin.Default()
router.POST("/", PostMethod)
  router.GET("/", GetMethod)
listenPort := "4000"
router.Run(":"+listenPort)
}
```
- Now, run this code to test if your server is working.


# Authentication

- Now that you’ve set up our basic web server, we can start adding in the elements for our authentication API.

- let's start with `func main()`:

```go
func main() {
  router := gin.Default()
subRouterAuthenticated := router.Group("/api/v1/PersonId", gin.BasicAuth(gin.Accounts{
    "admin": "adminpass",
  }))
listenPort := "1357"
router.Run(":"+listenPort)
}
```

- passing the query string parameters

```go
func main() {
  router := gin.Default()
subRouterAuthenticated := router.Group("/api/v1/PersonId", gin.BasicAuth(gin.Accounts{
    "admin": "adminpass",
  }))
subRouterAuthenticated.GET("/:IdValue", GetMethod)
listenPort := "1357"
  router.Run(":"+listenPort)
}
```

# Adjusting our GetMethod()

- It fetches and prints the Person `IdValue` from the query string parameter passed in the API URL:

```go

func GetMethod(c *gin.Context) {
  fmt.Println("\n'GetMethod' called")
  IdValue := c.Params.ByName("IdValue")
  message := "GetMethod Called With Param: " + IdValue
  c.JSON(http.StatusOK, message)
ReqPayload := make([]byte, 1024)
  ReqPayload, err := c.GetRawData()
  if err != nil {
        fmt.Println(err)
        return
  }
  fmt.Println("Request Payload Data: ", string(ReqPayload))
}
```
#

# For efficient practice check my `main.go` file

# Testing With ngrok

- First, run your Go application

```sh
go run main.go
```

1. Download [ngrok](https://dl.equinox.io/ngrok/ngrok/stable)
2. Extract the ngrok executable in a folder on your server.
3. Start ngrok on port `1357` (which is the port you selected in your code).
- p.s you can configure your port

```sh
./ngrok http 1357
```

- The result:

```sh
ngrok by @emmyc                                               (Ctrl+C to quit)
Session Status                online                                                         
Session Expires               7 hours, 12 minutes                                            
Version                       2.3.35                                                         
Region                        Netherlands (nl)                                             
Web Interface                 http://127.0.0.1:4040                                          
Forwarding                    http://ca6d2c4cee3e.ngrok.io -> http://localhost:4000          
Forwarding                    https://ca6d2c4cee3e.ngrok.io -> http://localhost:4000
```

- This will generate a random dynamic URL where you can test your API and log in with your login details to test if it works.

- After you log in, it will show:

```sh
GetMethod Called With Param: Id456
```

- Great It works!
