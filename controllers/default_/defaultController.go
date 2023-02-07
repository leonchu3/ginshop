package default_

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

// func (con DefaultController) Index(c *gin.Context) {
// 	c.SetCookie("username", "张三", 3600, "/", "localhost", false, true)
// 	c.HTML(http.StatusOK, "default/index.html", gin.H{
// 		"msg": "我是一个msg",
// 	})
// }

// func (con DefaultController) News(c *gin.Context) {

// 	username, _ := c.Cookie("username")
// 	c.String(http.StatusOK, "News--cookie="+username)
// }

func (con DefaultController) Index(c *gin.Context) {
	session := sessions.Default(c)
	session.Options(sessions.Options{
		MaxAge: 3600 * 6, //单位是秒 过期时间
	})
	session.Set("username", "张三111")
	session.Save()

	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg": "我是一个msg",
	})
}

func (con DefaultController) News(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")

	c.String(http.StatusOK, "username=%v", username)
}

func (con DefaultController) Shop(c *gin.Context) {

	username, _ := c.Cookie("username")
	c.String(http.StatusOK, "Shop--cookie="+username)
}

func (con DefaultController) DeleteCookie(c *gin.Context) {
	// 删除cookie
	c.SetCookie("username", "张三", -1, "/", "localhost", false, true)
	c.String(http.StatusOK, "删除成功！")
}
