package main

import (
  "fmt"
  "io"
  "log"
  "flag"
  "net/http"
  "github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
  message := "Ahoy! This gateway seems to be working fine!"
  c.JSON(http.StatusOK, gin.H{"message": message})
}

// TODO: add error handling and stuff
func miniMarketApiHealthCheck(c *gin.Context) {
  url := "http://localhost:3001/api/health_check"
  resp, err := http.Get(url)
  
  if err != nil { log.Fatal(err) }
  defer resp.Body.Close()

  bodyBytes, err := io.ReadAll(resp.Body)
  
  c.Data(resp.StatusCode, "application/json", []byte(bodyBytes))
}

// TODO:
// - add error handling of routes not found
// - add some way to config new api endpoints
func main() {
  var port string
  flag.StringVar(&port, "p", "3000", "Please specify a port number")
  flag.Parse()

  router := gin.Default()
  router.GET("/gateway/health_check", healthCheck)
  router.GET("/mini-market-api/api/health_check", miniMarketApiHealthCheck)

  fmt.Println("Ahoy! Server started!")
  router.Run("localhost:"+port)
}

