package itying

import (
	"ginshop/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	//1、获取顶部导航
	topNavList := []models.Nav{}
	models.DB.Where("status=1 AND position=1").Find(&topNavList)
	//2、获取轮播图数据
	focusList := []models.Focus{}
	models.DB.Where("status=1 AND focus_type=1").Find(&focusList)
	//3、获取分类的数据
	goodsCateList := []models.GoodsCate{}
	//https://gorm.io/zh_CN/docs/preload.html
	models.DB.Where("pid = 0 AND status=1").Order("sort DESC").Preload("GoodsCateItems", func(db *gorm.DB) *gorm.DB {
		return db.Where("goods_cate.status=1").Order("goods_cate.sort DESC")
	}).Find(&goodsCateList)

	// fmt.Println(focusList)
	//4, 获取中间导航
	middleNavList := []models.Nav{}
	models.DB.Where("status=1 AND position=2").Find(&middleNavList)

	for i := 0; i < len(middleNavList); i++ {
		relation := strings.ReplaceAll(middleNavList[i].Relation, "，", ",")
		relationIds := strings.Split(relation, ",")

		goodsList := []models.Goods{}
		models.DB.Where("id in ?", relationIds).Select("id,title,goods_img,price").Find(&goodsList)
		middleNavList[i].GoodsItems = goodsList
	}

	//手机
	phoneList := models.GetGoodsByCategory(35, "best", 8)

	//配件
	// otherList := models.GetGoodsByCategory(47, "best", 8)

	// c.JSON(http.StatusOK, gin.H{
	// 	"phoneList": phoneList,
	// })

	c.HTML(http.StatusOK, "itying/index/index.html", gin.H{
		"topNavList":    topNavList,
		"focusList":     focusList,
		"goodsCateList": goodsCateList,
		"middleNavList": middleNavList,
		"phoneList":     phoneList,
		// "otherList":     otherList,
	})

}

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

// func (con DefaultController) Index(c *gin.Context) {
// 	session := sessions.Default(c)
// 	session.Options(sessions.Options{
// 		MaxAge: 3600 * 6, //单位是秒 过期时间
// 	})
// 	session.Set("username", "张三111")
// 	session.Save()

// 	c.HTML(http.StatusOK, "default/index.html", gin.H{
// 		"msg": "我是一个msg",
// 	})
// }

// func (con DefaultController) Thumbnail1(c *gin.Context) {
// 	//按宽度进行比例缩放 输入输出都是文件
// 	filename := "static/upload/0.jpg"
// 	savepath := "static/upload/0_600.jpg"
// 	err := go_image.ScaleF2F(filename, savepath, 600)
// 	if err != nil {
// 		c.String(200, "生成图片失败")
// 		fmt.Println("---------------")
// 		fmt.Println(err)
// 		return
// 	}
// 	c.String(200, "生成图片成功")

// }

// func (con DefaultController) Thumbnail2(c *gin.Context) {
// 	filename := "static/upload/0.jpg"
// 	savepath := "static/upload/0_400.jpg"
// 	// 按宽度和高度进行比例缩放， 输入输出都是文件
// 	err := go_image.ThumbnailF2F(filename, savepath, 400, 200)
// 	if err != nil {
// 		c.String(200, "生成图片失败")
// 		return
// 	}
// 	c.String(200, "生成图片成功")

// }
// func (con DefaultController) Qrcode1(c *gin.Context) {
// 	var png []byte
// 	png, err := qrcode.Encode("love you, 娇娇 ❤❤", qrcode.Highest, 500)
// 	if err != nil {
// 		c.String(200, "生成二维码失败")
// 		return
// 	}
// 	c.String(200, string(png))
// }

// func (con DefaultController) Qrcode2(c *gin.Context) {
// 	savepath := "static/upload/qrcode.png"
// 	err := qrcode.WriteFile("你是我的小宝贝er~", qrcode.Medium, 556, savepath)
// 	if err != nil {
// 		c.String(200, "生成二维码失败")
// 		return
// 	}
// 	file, _ := ioutil.ReadFile(savepath)
// 	c.String(200, string(file))
// }

// func (con DefaultController) News(c *gin.Context) {
// 	session := sessions.Default(c)
// 	username := session.Get("username")

// 	c.String(http.StatusOK, "username=%v", username)
// }

// func (con DefaultController) Shop(c *gin.Context) {

// 	username, _ := c.Cookie("username")
// 	c.String(http.StatusOK, "Shop--cookie="+username)
// }

// func (con DefaultController) DeleteCookie(c *gin.Context) {
// 	// 删除cookie
// 	c.SetCookie("username", "张三", -1, "/", "localhost", false, true)
// 	c.String(http.StatusOK, "删除成功！")
// }
