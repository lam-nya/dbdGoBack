package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  router.GET("/", func(context *gin.Context) {
    context.JSON(http.StatusOK, gin.H{"data": "hello world"})
  })

  router.GET("/auth", func(context *gin.Context) {
    context.JSON(200, gin.H{"username": "username1"})
    context.JSON(http.StatusOK, gin.H{"userid": "userid"})
  })

  err := router.Run(":81")
  if err != nil {
    panic("[Error] failled to start GIN:"+err.Error())
    return
  }
}
