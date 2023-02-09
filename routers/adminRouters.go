package routers

import (
	"ginshop/controllers/admin"
	"ginshop/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {

	adminRouters := r.Group("/admin", middlewares.InitAdminAuthMiddleware)

	adminRouters.GET("/", admin.MainController{}.Index)
	adminRouters.GET("/welcome", admin.MainController{}.Welcome)

	adminRouters.GET("/login", admin.LoginController{}.Index)
	adminRouters.POST("/doLogin", admin.LoginController{}.Dologin)
	adminRouters.GET("/loginOut", admin.LoginController{}.Loginout) //get和post请求一定要注意别用错

	adminRouters.GET("/captcha", admin.LoginController{}.Captcha)
	// adminRouters.GET("/doLogin", admin.LoginController{}.VerifyCaptcha)

	adminRouters.GET("/manager", admin.ManagerController{}.Index)
	adminRouters.GET("/manager/add", admin.ManagerController{}.Add)
	adminRouters.POST("/manager/doAdd", admin.ManagerController{}.DoAdd)
	adminRouters.GET("/manager/edit", admin.ManagerController{}.Edit)
	adminRouters.POST("/manager/doEdit", admin.ManagerController{}.DoEdit)
	adminRouters.GET("/manager/delete", admin.ManagerController{}.Delete)

	adminRouters.GET("/focus", admin.FocusController{}.Index)
	adminRouters.GET("/focus/add", admin.FocusController{}.Add)
	adminRouters.GET("/focus/edit", admin.FocusController{}.Edit)
	adminRouters.GET("/focus/delete", admin.FocusController{}.Delete)

	adminRouters.GET("/role", admin.RoleController{}.Index)
	adminRouters.GET("/role/add", admin.RoleController{}.Add)
	adminRouters.GET("/role/edit", admin.RoleController{}.Edit)
	adminRouters.GET("/role/delete", admin.RoleController{}.Delete)
	adminRouters.POST("/role/doAdd", admin.RoleController{}.DoAdd)
	adminRouters.POST("/role/doEdit", admin.RoleController{}.DoEdit)

	adminRouters.GET("/access", admin.AccessController{}.Index)
	adminRouters.GET("/access/add", admin.AccessController{}.Add)
	adminRouters.POST("/access/doAdd", admin.AccessController{}.DoAdd)
	adminRouters.GET("/access/edit", admin.AccessController{}.Edit)
	adminRouters.POST("/access/doEdit", admin.AccessController{}.DoEdit)
	adminRouters.GET("/access/delete", admin.AccessController{}.Delete)
}
