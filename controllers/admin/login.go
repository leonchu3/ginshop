package admin

import (
	"encoding/json"
	"fmt"
	"ginshop/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	c.HTML(200, "admin/login/login.html", gin.H{})
}

func (con LoginController) Dologin(c *gin.Context) {
	captchaId := c.PostForm("captchaId")
	username := c.PostForm("username")
	password := c.PostForm("password")

	captchaValue := c.PostForm("verifyValue")
	if flag := models.VerifyCaptcha(captchaId, captchaValue); flag == true {
		userinfoList := []models.Manager{}
		password := models.Md5(password)
		models.DB.Where("username=? AND password=?", username, password).Find(&userinfoList)

		if len(userinfoList) > 0 {
			session := sessions.Default(c)

			//session.set没办法直接保存切片 把结构体转换成json字符串
			userifnoByte, _ := json.Marshal(userinfoList)
			session.Set("userinfo", string(userifnoByte))
			session.Save()
			con.Success(c, "登录成功", "/admin")
		} else {
			con.Error(c, "用户名或密码错误", "/admin/login")
		}
	} else {
		con.Error(c, "验证码验证失败", "/admin/login")
	}

}

func (con LoginController) Captcha(c *gin.Context) {
	id, b64s, err := models.MakeCaptcha()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}

func (con LoginController) Loginout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userinfo")
	session.Save()
	con.Success(c, "退出登录成功", "/admin/login")
}

func (con LoginController) VerifyCaptcha(c *gin.Context) {
	c.String(200, "dologin")
}
