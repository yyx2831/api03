package router

import (
	"api03/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRegister := Router.Group("user")
	{
		UserRegister.GET("list", api.GetUserList)
	}
	//Router.POST("/user/login", UserLogin)
	//Router.POST("/user/register", UserRegister)
	//Router.GET("/user/info", UserInfo)
	//Router.GET("/user/logout", UserLogout)
}
