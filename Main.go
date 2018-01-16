package main

import (
	"./login"
	"github.com/gin-gonic/gin"
)

func main() {

	c := gin.Default()

	v1 := c.Group("/v1")

	{
		v1.GET("/login", func(c *gin.Context) {
			login.Login(c)
		})

		v1.GET("/logout", func(c *gin.Context) {
			login.Logout(c)
		})

		v1.GET("/register", func(c *gin.Context) {
			login.Register(c)
		})
	}


	v2 := c.Group("/v2")

	{
		v2.POST("/login", func(c *gin.Context) {
			login.Login(c)
		})

		v2.POST("/logout", func(c *gin.Context) {
			login.Logout(c)
		})
	}


	c.Run(":8199")


}
