package main

import (
	// git@github.com:go-playground/locales.git
	// git@github.com:ugorji/go/codec.git
	// git@github.com:astaxie/beego.git

	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/ping", func(c *gin.Context) {
		//c.String(http.StatusOK, "hello World!")
		c.String(http.StatusOK,"")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	err := r.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
