package main
import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)
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
func main() {
  router := gin.Default()
}
