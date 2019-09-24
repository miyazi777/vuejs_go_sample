package main

import "github.com/gin-gonic/gin"

func test1(c *gin.Context) {
	c.String(200, "Hello World2")
}

func test2(c *gin.Context) {
	c.String(200, "fuga")
}

func main() {
	r := gin.Default()
	r.GET("/", test1)
	r.GET("/hoge", test2)

	r.Run(":3000")
}
