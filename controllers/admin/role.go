package admin

import (
	"ginshop/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	BaseController
}

func (con RoleController) Index(c *gin.Context) {

	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})
}

func (con RoleController) Add(c *gin.Context) {
	c.HTML(200, "admin/role/add.html", gin.H{})
}

func (con RoleController) DoAdd(c *gin.Context) {

	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	if title == "" {
		con.Error(c, "角色的标题不能为空", "/admin/role/add")
	} else {
		role := models.Role{}
		role.Title = title
		role.Description = description
		role.Status = 1
		role.AddTime = int(models.GetUnix())

		err := models.DB.Create(&role).Error
		if err != nil {
			con.Error(c, "添加角色失败，请重试！", "/admin/role/add")
		} else {
			con.Success(c, "添加角色成功", "/admin/role")
		}
	}

}

func (con RoleController) DoEdit(c *gin.Context) {
	id, err := models.Int(c.PostForm("id")) //获得的是提交表单的参数
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
	} else {
		title := strings.Trim(c.PostForm("title"), " ")
		description := strings.Trim(c.PostForm("description"), " ")
		if title == "" {
			con.Error(c, "角色的标题不能为空", "/admin/role/edit")
		}
		role := models.Role{Id: id}
		models.DB.Find(&role)
		role.Title = title
		role.Description = description

		err2 := models.DB.Save(&role).Error
		if err2 != nil {
			con.Error(c, "修改数据失败", "/admin/role/edit?id="+models.String(id))
		} else {
			con.Success(c, "修改数据成功", "/admin/role/edit?id="+models.String(id))
		}
	}
}

func (con RoleController) Edit(c *gin.Context) {

	id, err := models.Int(c.Query("id")) //query获得的是url里的参数
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
	} else {
		role := models.Role{Id: id}
		models.DB.Find(&role)
		c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
			"role": role,
		})
	}
}
func (con RoleController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "删除数据错误", "/admin/role")
	} else {
		role := models.Role{Id: id}
		models.DB.Delete(&role)
		con.Success(c, "删除数据成功", "/admin/role")
	}
}
