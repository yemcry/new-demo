package mysql

import (
	"fmt"
	"server/models"
	"time"
)

func UserList() []models.User {
	db := GetDB()
	var user []models.User
	db.Find(&user)
	return user
}

func UserInfo(name string) models.User {
	db := GetDB()
	var user models.User
	db.Where("username = ?", name).First(&user)
	return user
}

func UserLogin(u models.User) error {
	db := GetDB()
	if err := db.Where("username = ? AND password = ?",
		u.Username, u.Password).First(&u).Error; err != nil {
		return err
	}
	return nil
}

func UserRegister(u models.User) error {
	db := GetDB()
	user := models.User{
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		AvatarURL: "10.jpg",
		Status:    1,
	}
	if err := db.Create(&user).Error; err != nil {
		fmt.Printf("%s用户注册失败！！！", u.Username)
		return err
	}
	return nil
}

func UserHomi(name string) []models.FriendInfo {
	db := GetDB()
	var user models.User
	// 查找用户
	db.Where("username=?", name).First(&user)

	var friendships []models.FriendShip
	// 查找用户发起的好友关系
	db.Where("user_id = ? AND status = ?", user.ID, 2).Find(&friendships)

	var ships2 []models.FriendShip
	// 查找好友发起的关系
	db.Where("friend_id = ? AND status = ?", user.ID, 2).Find(&ships2)

	// 创建切片存储最终的好友信息
	var friendsInfo []models.FriendInfo

	// 处理发起的关系
	for _, f := range friendships {
		friend := models.FriendInfo{
			FriendID:   f.FriendID,
			FriendName: f.Friendname,
			Status:     f.Status,
		}
		// 获取好友头像
		var friendUser models.User
		db.Where("user_id = ?", f.FriendID).First(&friendUser)
		friend.AvatarURL = friendUser.AvatarURL

		friendsInfo = append(friendsInfo, friend)
	}

	// 处理被发起的关系
	for _, f := range ships2 {
		friend := models.FriendInfo{
			FriendID:   f.UserID,
			FriendName: f.Username,
			Status:     f.Status,
		}
		// 获取好友头像
		var friendUser models.User
		db.Where("user_id = ?", f.UserID).First(&friendUser)
		friend.AvatarURL = friendUser.AvatarURL

		friendsInfo = append(friendsInfo, friend)
	}
	return friendsInfo
}

func UserApply(name string) []models.FriendShip {
	db := GetDB()
	var ship []models.FriendShip
	db.Where("friendname = ? AND status = ?", name, 1).Find(&ship)
	return ship
}

// 同意好友
func UserAgree(info models.AgreeInfo) error {
	db := GetDB()
	var ship models.FriendShip
	db.Where("username=? AND friendname=?", info.Friendname, info.Username).First(&ship)
	ship.Status = 2
	if err := db.Save(&ship).Error; err != nil {
		return err
	}
	return nil
}
