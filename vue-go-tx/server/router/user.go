package router

import (
	"server/services"

	"github.com/gin-gonic/gin"
)

func UserLogin(r *gin.Engine) {
	r.GET("/users", services.UserList)     // 查询所有用户
	r.GET("/user_info", services.UserInfo) //返回当前用户信息
	r.GET("/homi", services.UserHomi)      //返回当前用户好友
	r.GET("/apply", services.UserApply)    //返回当前用户申请列表
	r.POST("/agree", services.UserAgree)   // 同意好友
	r.POST("/login", services.Login)       // 登录
	r.POST("/register", services.Register) // 注册
}
