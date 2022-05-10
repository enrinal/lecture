package main

import "github.com/gin-gonic/gin"

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type SuccessResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    LoginRequest `json:"data"`
}

func LoginHandler(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, ErrorResponse{
			Code:    400,
			Message: "Bad Request",
			Error:   err.Error(),
		})
		return
	}

	if request.Username != "admin" || request.Password != "admin" {
		c.JSON(401, ErrorResponse{
			Code:    401,
			Message: "Unauthorized",
			Error:   "Unauthorized",
		})
		return
	}

	c.JSON(200, SuccessResponse{
		Code:    200,
		Message: "Success",
		Data:    request,
	})

}

func main() {
	r := gin.Default()

	r.POST("/login", LoginHandler)

	r.Run(":8080")
}
