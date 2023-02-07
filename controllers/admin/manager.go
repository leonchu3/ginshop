package admin

import "github.com/gin-gonic/gin"

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {
	c.HTML(200, "admin/manager/index.html", gin.H{})
}

func (con ManagerController) Add(c *gin.Context) {
	c.HTML(200, "admin/manager/add.html", gin.H{})
}

func (con ManagerController) Edit(c *gin.Context) {
	c.HTML(200, "admin/manager/edit.html", gin.H{})
}

func (con ManagerController) Delete(c *gin.Context) {
	c.String(200, "admin/manager/delete.html", gin.H{})
}
