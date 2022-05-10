package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, HelloResponse{
		Message: "Hello, World!",
		Code:    http.StatusOK,
	})
}

func helloStudent(c *gin.Context) {
	name := c.Param("name") // get name from url param
	c.JSON(http.StatusOK, HelloResponse{
		Message: "Hello, " + name,
		Code:    http.StatusOK,
	})
}

func helloStudentHobby(c *gin.Context) {
	name := c.Param("name")   // get name from url param
	hobby := c.Param("hobby") // get hobby from url param

	c.JSON(http.StatusOK, HelloResponse{
		Message: "Hello, " + name + " and your hobby is " + hobby,
		Code:    http.StatusOK,
	})
}

func helloTeacher(c *gin.Context) {
	teacherName := c.Query("name") // get teacher name from url param
	teacherName = c.DefaultQuery("name", "admin")
	c.JSON(http.StatusOK, HelloResponse{
		Message: "Hello, " + teacherName,
		Code:    http.StatusOK,
	})
}

func main() {
	r := gin.Default()

	r.GET("/sample", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	r.GET("/hello", helloHandler)
	r.GET("/hello/:name", helloStudent) // set name to url param
	r.GET("/hello/:name/*hobby", helloStudentHobby)
	r.GET("/hello-teacher", helloTeacher)

	r.Run(":8080")
}
