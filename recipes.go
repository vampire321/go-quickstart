package main

import "github.com/gin-gonic/gin"

func recipes(){
	r := gin.Default()
	r.GET("/recipes",func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"pong",
		})
	})
}