package service

import (
	"fmt"

	"Backend/service/config"

	"github.com/gin-gonic/gin"
)

func RunApis() {
	// 初始化DI
	fmt.Println("*****Web Service Start*****")
	fmt.Printf("\n")
	_getPort := config.Load().GinConfig.Port
	fmt.Println("get config port: ", _getPort)

	// gin.SetMode(gin.DebugMode) // 切換到 Release 模式
	_gin := gin.Default() //宣告_gin元件

	_gin.GET("/ping", ping) //Http ping 指令

	// _gin.Run(":8080") //執行WebService part:8080
	_gin.Run(_getPort)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, world",
	})

}
