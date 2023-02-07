package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddlewares(c *gin.Context) {
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL) 
	c.Set("username", "张三") //设置中间件信息用于共享

	//定义一个goroutine统计日志  当中间件或handler中启动新的goroutine时要使用c的副本

	cCp := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done in path " + cCp.Request.URL.Path)
	}()
}
