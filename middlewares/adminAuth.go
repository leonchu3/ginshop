package middlewares

import (
	"encoding/json"
	"fmt"
	"ginshop/models"
	"os"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

func judgeRedirect(c *gin.Context, pathname string) {
	if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
		c.Redirect(302, "/admin/login")
	}
}

func excludeAuthPath(urlPath string) bool {
	config, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	excludeAuthPath := config.Section("").Key("excludeAuthPath").String()
	excludeAuthPathSlice := strings.Split(excludeAuthPath, ",")
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
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
		// err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)

		if err != nil || !(len(userinfoStruct) > 0 && userinfoStruct[0].Username != "") {
			judgeRedirect(c, pathname)
		} else {
			urlPath := strings.Replace(pathname, "/admin/", "", 1)
			if userinfoStruct[0].IsSuper == 0 && !excludeAuthPath("/"+urlPath) {
				//1 根据角色获取当前角色的权限列表 然后把权限id放在一个map类型的对象里面

				roleAccess := []models.RoleAccess{}
				models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
				roleAccessMap := make(map[int]int)
				for _, v := range roleAccess {
					roleAccessMap[v.AccessId] = v.AccessId
				}
				//2 获取当前访问的url对应的权限id 判断权限id是否在角色对应的权限
				access := models.Access{}
				models.DB.Where("url=?", urlPath).Find(&access)

				if _, ok := roleAccessMap[access.Id]; !ok {
					c.String(200, "没有权限")
					c.Abort()
				}
			}
		}

	} else {
		//用户没有登录
		//排除不需要做权限判断的路由
		judgeRedirect(c, pathname)

	}
}
