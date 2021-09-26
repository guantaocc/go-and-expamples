package main

import (
	"frame/gee"
	"net/http"
)

func main() {
	e := gee.New()
	e.GET("/", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"code":    0,
			"message": "gin go run",
		})
	})

	e.Run(":8080")
}
