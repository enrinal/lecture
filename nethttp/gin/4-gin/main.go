package main

import "github.com/gin-gonic/gin"

func helloHandler(c *gin.Context) {
	c.String(200, "Hello World")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/hello", helloHandler)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
