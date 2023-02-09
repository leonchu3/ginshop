package admin

import (
	"fmt"
	"ginshop/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {

	mangeerList := []models.Manager{}
	models.DB.Preload("Role").Find(&mangeerList)

	// c.JSON(http.StatusOK, gin.H{
	// 	"managerList": mangeerList,
	// })
	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": mangeerList,
	})
}

func (con ManagerController) Add(c *gin.Context) {
	//获取所有的角色
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}

func (con ManagerController) DoAdd(c *gin.Context) {
	roleId, err1 := models.Int(c.PostForm("role_id"))

	if err1 != nil {
		con.Error(c, "传入数据错误", "admin/manager/add")
	}

	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	if len(username) < 2 || len(password) < 6 {
		con.Error(c, "用户或者密码的长度不合法", "/admin/manager/add")
	} else {
		managerList := []models.Manager{}
		models.DB.Where("username=?", username).Find(&managerList)
		if len(managerList) > 0 {
			con.Error(c, "此管理员已存在", "/admin/manager/add")
		} else {
			manager := models.Manager{
				Username: username,
				Password: models.Md5(password),
				RoleId:   roleId,
				Status:   1,
				AddTime:  int(models.GetUnix()),
				Mobile:   mobile,
				Email:    email,
			}
			err2 := models.DB.Create(&manager).Error
			if err2 != nil {
				con.Error(c, "增加管理员失败", "/admin/manager/add")
			} else {
				con.Success(c, "增加管理员成功", "/admin/manager")
			}
		}
	}
}

func (con ManagerController) Edit(c *gin.Context) {
	managerId, err := models.Int(c.Query("id"))

	if err != nil {
		con.Error(c, "传入数据错误", "admin/manager")
	} else {
		manager := models.Manager{Id: managerId}
		models.DB.Find(&manager)

		roleList := []models.Role{}
		models.DB.Find(&roleList)

		fmt.Println(roleList)
		c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
			"manager":  manager,
			"roleList": roleList,
		})
	}
}

func (con ManagerController) DoEdit(c *gin.Context) {
	id, err1 := models.Int(c.PostForm("id"))

	if err1 != nil {
		con.Error(c, "传入数据错误", "admin/manager")
		return
	}

	roleId, err2 := models.Int(c.PostForm("role_id"))

	if err2 != nil {
		con.Error(c, "传入数据错误", "admin/manager")
		return
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	email := strings.Trim(c.PostForm("email"), " ")

	//执行修改
	manager := models.Manager{Id: id}
	models.DB.Find(&manager)

	manager.Username = username
	manager.Mobile = mobile
	manager.Email = email
	manager.RoleId = roleId

	if password != "" {
		if len(username) < 2 || len(password) < 6 {
			con.Error(c, "用户或者密码的长度不合法 密码长度不小于6位", "/admin/manager/edit?id="+models.String(id))
		} else {
			manager.Password = models.Md5(password)
		}
	}
	err3 := models.DB.Save(&manager).Error
	if err3 != nil {
		con.Error(c, "修改数据失败", "/admin/manager/edit?id="+models.String(id))
	} else {
		con.Success(c, "修改数据成功", "/admin/manager")
	}

}

func (con ManagerController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "删除管理员错误", "/admin/manager")
	} else {
		manager := models.Manager{Id: id}
		models.DB.Delete(&manager)
		con.Success(c, "删除管理员成功", "/admin/manager")
	}
}
