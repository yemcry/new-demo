package services

import (
	"fmt"
	"net/http"
	"server/models"
	"server/models/common/response"
	"server/mysql"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	res := mysql.UserList()
	c.JSON(http.StatusOK, res)
}

func UserInfo(c *gin.Context) {
	name := c.Query("username")
	user := mysql.UserInfo(name)
	response.OkWithData(user, c)
}

func Login(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)
	if err := mysql.UserLogin(user); err != nil {
		fmt.Println("登录失败")
		response.Fail(c)
	} else {
		response.Ok(c)
	}
}

func Register(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)
	if err := mysql.UserRegister(user); err != nil {
		fmt.Println("注册失败")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "注册失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    "注册成功",
		})
	}
}

// 返回当前用户好友
func UserHomi(c *gin.Context) {
	name := c.Query("username")
	user := mysql.UserHomi(name)
	response.OkWithData(user, c)
}

// 返回当前用户申请列表
func UserApply(c *gin.Context) {
	name := c.Query("username")
	user := mysql.UserApply(name)
	response.OkWithData(user, c)
}

// 同意好友
func UserAgree(c *gin.Context) {
	var info models.AgreeInfo
	c.ShouldBind(&info)
	if err := mysql.UserAgree(info); err != nil {
		response.Fail(c)
	} else {
		response.OkWithData("ok", c)
	}

}
