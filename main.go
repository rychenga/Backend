// demo project main.go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("*****Web Service Start*****")
	fmt.Printf("\n")

	_gin := gin.Default() //宣告_gin元件

	_gin.GET("/ping", ping) //Http ping 指令

	_gin.Run(":8080") //執行WebService part:8080

}

// HTTP ping 指令
func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, world",
	})

}
