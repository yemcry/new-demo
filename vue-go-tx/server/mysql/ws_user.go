package mysql

import (
	"fmt"
	"server/models"
	"time"

	"gorm.io/gorm"
)

func UserStatus(name string, s int) {
	db := GetDB()
	db.Model(models.User{}).Where("username=?", name).Update("status", s)
}

func GetUserID(uname string, fname string) (uuid uint, fuid uint, err error) {
	db := GetDB()
	var user models.User
	var fuser models.User
	if err = db.Where("username=?", uname).First(&user).Error; err != nil {
		return 0, 0, err
	}
	if err = db.Where("username=?", fname).First(&fuser).Error; err != nil {
		return 0, 0, err
	}

	return user.ID, fuser.ID, nil
}

func AddFriends(data models.Data) (int, error) {
	db := GetDB()
	var user models.User
	if err := db.Where("username=?", data.Friendname).First(&user).Error; err != nil {
		fmt.Println("没有这个人！！！")
		return 0, err
	}

	// 调用 CheckFriendship 函数
	result, err := CheckFriendship(db, data.Username, data.Friendname)
	if err != nil {
		return 0, err
	}

	if result == 1 {
		fmt.Println("您已向此用户发送好友申请！")
		return 1, nil
	}

	// 如果没有申请记录，可以创建新的好友申请
	uuid, fuid, _ := GetUserID(data.Username, data.Friendname)
	f := models.FriendShip{
		UserID:     uuid,
		FriendID:   fuid,
		Username:   data.Username,
		Friendname: data.Friendname,
		Remark:     data.Remark,
		Status:     1, // 状态设置为待处理
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Create(&f).Error; err != nil {
		fmt.Println("添加好友失败")
		return 0, err
	}
	fmt.Println("添加好友成功")
	return 0, nil
}

func CheckFriendship(db *gorm.DB, username, friendname string) (int, error) {
	var friendship models.FriendShip

	// 检查当前用户向对方的申请
	if err := db.Where("username = ? AND friendname = ?", username, friendname).First(&friendship).Error; err == nil {
		return 1, nil // 返回 1 重复申请
	}

	// 检查对方是否向当前用户申请
	if err := db.Where("username = ? AND friendname = ?", friendname, username).First(&friendship).Error; err == nil {
		return 1, nil
	}

	return 0, nil // 返回 0 正常申请
}
