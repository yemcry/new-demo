package models

// User 表
type User struct {
	ID        uint   `json:"user_id" gorm:"column:user_id"` // 用户ID
	Username  string `json:"username"`                      // 用户名
	Password  string `json:"password"`                      // 密码
	AvatarURL string `json:"avatar_url"`                    // 头像URL
	Status    uint   `json:"status"`                        // 用户状态	1.离线 2.在线
	CreatedAt string `json:"created_at"`                    // 创建时间
}

// Friendship 表
type FriendShip struct {
	ID         uint   `json:"id"`         // 关系ID
	UserID     uint   `json:"user_id"`    // 用户ID
	FriendID   uint   `json:"friend_id"`  // 好友ID
	Username   string `json:"username"`   // 用户名
	Friendname string `json:"friendname"` // 好友名
	Remark     string `json:"remark"`     //备注
	Status     uint   `json:"status"`     // 关系状态	1.未添加 2.已添加
	CreatedAt  string `json:"created_at"` // 创建时间
}

// Group 表
type Group struct {
	ID        uint   `json:"group_id"`     // 群组ID
	Name      string `json:"group_name"`   // 群组名称
	Avatar    string `json:"group_avatar"` // 群组头像
	Depict    string `json:"group_depict"` // 群组描述
	CreatedAt string `json:"created_at"`   // 创建时间
	OwnerID   uint   `json:"owner_id"`     // 群主ID
	OwnerName string `json:"owner_name"`   // 群主ID
}

// GroupMember 表
type GroupMember struct {
	ID       uint   `json:"id"`        // 成员ID
	GroupID  uint   `json:"group_id"`  // 群组ID
	UserID   uint   `json:"user_id"`   // 用户ID
	Username string `json:"username"`  // 用户ID
	Status   uint   `json:"status"`    // 用户ID
	JoinedAt string `json:"joined_at"` // 加入时间
	Role     string `json:"role"`      // 用户角色
}

// Message 表
type Message struct {
	ID         uint   `json:"id"`          // 消息ID
	ChatID     uint   `json:"chat_id"`     // 聊天ID
	SenderID   uint   `json:"sender_id"`   // 发送者ID
	SenderName uint   `json:"sender_name"` // 发送者ID
	Content    string `json:"content"`     // 消息内容
	Timestamp  string `json:"timestamp"`   // 消息时间戳
	Type       string `json:"type"`        // 消息类型
	ChatType   string `json:"chat_type"`   // 聊天类型
}

type AgreeInfo struct {
	Username   string `json:"username"`
	Friendname string `json:"friendname"`
}

type FriendInfo struct {
	FriendID   uint   `json:"friend_id"`   // 好友ID
	FriendName string `json:"friend_name"` // 好友名
	AvatarURL  string `json:"avatar_url"`  // 好友头像URL
	Status     uint   `json:"status"`      // 关系状态
}
