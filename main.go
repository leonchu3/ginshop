package main

import (
	"fmt"
	"os"

	"ginshop/routers"
	"text/template"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

// 时间戳转换为日期
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func Println(str1 string, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + "----" + str2
}

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" form:"content"`
}

func initMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()

	fmt.Println("1-我是一个中间件")
	//调用该请求的剩余处理程序
	c.Next()
	// c.Abort()

	fmt.Println("2-我是一个中间件")

	end := time.Now().UnixNano()
	fmt.Println(end - start)
}

func main() {
	r := gin.Default()

	// 自定义模板函数  注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
		"Println":    Println,
	})

	r.LoadHTMLGlob("templates/**/**/*")

	// 配置静态web服务
	r.Static("/static", "./static")

	//配置session中间件
	// store := cookie.NewStore([]byte("secret11111"))
	// // 设置 session 中间件，参数 mysession，指的是 session 的名字，也是 cookie 的名字
	// // store 是前面创建的存储引擎，我们可以替换成其他存储引擎
	// r.Use(sessions.Sessions("mysession", store))

	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)

	//演示go.ini 的使用
	config, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	fmt.Println(config.Section("").Key("app_name").String())
	fmt.Println(config.Section("redis").Key("ip").String())
	fmt.Println(config.Section("mysql").Key("password").String())

	//给ini写入数据
	config.Section("").Key("app_name").SetValue("哈哈 gin")
	config.Section("").Key("admin_path").SetValue("nihao/gin")
	config.SaveTo("./conf/app.ini")

	//前台
	//全局中间件
	// r.Use(initMiddleware)

	// // GET请求传值
	// r.GET("/", func(c *gin.Context) {
	// 	fmt.Println("这是一个首页")
	// 	// time.Sleep(time.Second)
	// 	c.String(200, "这是主页")
	// })

	// r.GET("/news", func(c *gin.Context) {
	// 	c.String(200, "新闻页面")
	// })

	r.Run()
}
