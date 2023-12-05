package testbind

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Bind struct {
}

func (b *Bind) Test(c *gin.Context) {
	data := make(map[string]string)
	if err := c.BindJSON(&data); err != nil {
		fmt.Printf("Error bind requestData: %v\n ", err)
		fmt.Fprintf(c.Writer, "Failed! Error bind requestData: %v ", err)
		return
	}
	podLabel := map[string]string{
		"yaml":    data["yaml"],
		"sender":  data["sender"],
		"user-id": data["user-id"],
	}
	fmt.Println(podLabel)
}
