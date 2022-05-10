package main

import "github.com/gin-gonic/gin"

func handleHelloV1(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World v1",
	})
}

func handleHelloV2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World v2",
	})
}

func handleHelloTeacherV1(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "Hello " + name + " v1",
	})
}

func handleHelloTeacherV2(c *gin.Context) {
	name := c.Query("name")
	c.JSON(200, gin.H{
		"message": "Hello " + name + " v2",
	})
}
func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello", handleHelloV1)
		v1.GET("/teacher/:name", handleHelloTeacherV1)
	}

	v2 := r.Group("/api/v2")
	{
		v2.GET("/hello", handleHelloV2)
		v2.GET("/teacher", handleHelloTeacherV2)
	}

	r.Run(":8080")
}
