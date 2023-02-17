package routers

import (
	"ginshop/controllers/itying"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {

	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", itying.DefaultController{}.Index)

		// defaultRouters.GET("/news", default_test.DefaultController{}.News)
		// defaultRouters.GET("/shop", default_test.DefaultController{}.Shop)

		// defaultRouters.GET("/thumbnail1", default_test.DefaultController{}.Thumbnail1)
		// defaultRouters.GET("/thumbnail2", default_test.DefaultController{}.Thumbnail2)
		// defaultRouters.GET("/qrcode1", default_test.DefaultController{}.Qrcode1)
		// defaultRouters.GET("/qrcode2", default_test.DefaultController{}.Qrcode2)

	}
}
