package utils

import "github.com/gin-gonic/gin"

type Response struct{
	Success bool `json:"success"`
	Message string `json:"message",omitempty`
	Data any `json:"data",omitempty`
	Error any `json:"error",omitempty`
}

func OK(ctx *gin.Context, statusCode int, data any){
	ctx.JSON(statusCode, Response{Success: true, Data: data})
}

func Fail(ctx *gin.Context, statusCode int, data any){
	ctx.JSON(statusCode, Response{Success: false, Data: data})
}
