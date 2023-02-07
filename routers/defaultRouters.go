package routers

import (
	"ginshop/controllers/default_"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {

	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", default_.DefaultController{}.Index)

		defaultRouters.GET("/news", default_.DefaultController{}.News)
		defaultRouters.GET("/shop", default_.DefaultController{}.Shop)
		defaultRouters.GET("/deletecookie", default_.DefaultController{}.DeleteCookie)
	}
}
