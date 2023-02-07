package admin

import "github.com/gin-gonic/gin"

type FocusController struct {
	BaseController
}

func (con FocusController) Index(c *gin.Context) {
	c.HTML(200, "admin/focus/index.html", gin.H{})
}

func (con FocusController) Add(c *gin.Context) {
	c.HTML(200, "admin/focus/add.html", gin.H{})
}

func (con FocusController) Edit(c *gin.Context) {
	c.HTML(200, "admin/focus/edit.html", gin.H{})
}

func (con FocusController) Delete(c *gin.Context) {
	c.String(200, "admin/focus/delete.html", gin.H{})
}
