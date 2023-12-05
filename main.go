package main

import (
	"fmt"
	"test/testbind"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("init func !")
}

func main() {
	r := gin.Default()

	tb := testbind.New(map[string]string{})
	cluster := r.Group("/test/testbind/bind") // Cluster
	{
		cluster.POST("/",
			func(c *gin.Context) {
				tb.Bind.Test(c)
			})
	}
	r.Run(":9999")
}
