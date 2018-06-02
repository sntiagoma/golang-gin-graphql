package main

import (
	"github.com/gin-gonic/gin"
)

var DB = make(map[string]string)

func setupRouter() *gin.Engine {
	gin.DisableConsoleColor()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pongx")
	})
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := DB[user]
		if ok {
			c.JSON(200, gin.H{"user": user, "value": value})
		} else {
			c.JSON(200, gin.H{"user": user, "status": "no value"})
		}
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar",
		"manu": "123",
	}))

	authorized.POST("admin", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r
}
