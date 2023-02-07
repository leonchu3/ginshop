package middlewares

import (
	"encoding/json"
	"fmt"
	"ginshop/models"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func judgeRedirect(c *gin.Context, pathname string) {
	if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
		c.Redirect(302, "/admin/login")
	}
}

func InitAdminAuthMiddleware(c *gin.Context) {
	// //获取userinfo对应的session
	// sessions := sessions.Default(c)
	// userinfo := sessions.Get("userinfo")

	// //类型断言 来判断 userinfo 是不是一个string
	// userinfoStr, ok := userinfo.(string)
	// if ok {
	// 	var userinfoStruct []models.Manager
	// 	json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
	// 	fmt.Println(userinfoStruct)
	// 	// c.HTML(http.StatusOK, "admin/main/index.html", gin.H{})
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"username": userinfoStruct[0].Username,
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"username": "session不存在",
	// 	})
	// }

	//1获取所有需要验证的url访问地址
	fmt.Println("InitAdminAuthMiddleware")
	pathname := strings.Split(c.Request.URL.String(), "?")[0]

	// 2获取Session里面保存的用户信息
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	//使用断言来判断userinfo是不是一个string
	userinfoStr, ok := userinfo.(string)

	if ok {
		//判断userinfo里面的信息是否存在
		var userinfoStruct []models.Manager
		err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)

		if err != nil || !(len(userinfoStruct) > 0 && userinfoStruct[0].Username != "") {
			judgeRedirect(c, pathname)
		}

	} else {
		//用户没有登录
		//排除不需要做权限判断的路由
		judgeRedirect(c, pathname)

	}
}
