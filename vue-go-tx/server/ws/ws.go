package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/models"
	"server/mysql"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有跨域请求
	},
}

var (
	systemUsername = "系统"
	mu             sync.Mutex
	clients        = make(map[string]*websocket.Conn)
) // 互斥锁

func HandleWebSocket(c *gin.Context) {
	username := c.Query("username") // 获取用户名
	if username == "" {
		fmt.Println("用户名为空")
		c.String(http.StatusBadRequest, "用户名不能为空")
		return
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("时间：%s\n用户：%s  连接了服务器\n", timestamp, username)
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("无法升级连接:", err)
		return
	}
	defer conn.Close()

	addClient(username, conn)
	defer removeClient(username)
	// 用户进入提示
	Hello := fmt.Sprintf("欢迎%s用户来到webSocket初始版！！！", username)
	BroadcastMessage(models.User{}, Hello, timestamp)
	for {
		if err := handleClientMessages(username, conn); err != nil {
			fmt.Println("客户端断开连接:", err)
			break
		}
	}
}

// 添加客户端
func addClient(username string, conn *websocket.Conn) {
	mu.Lock()
	clients[username] = conn
	mu.Unlock()
	mysql.UserStatus(username, 2)
	//sendOnlineUsers() // 每次添加用户后发送在线用户列表
}

// 移除客户端
func removeClient(username string) {
	mu.Lock()
	delete(clients, username)
	mu.Unlock()
	mysql.UserStatus(username, 1)
	//sendOnlineUsers() // 每次移除用户后发送在线用户列表
}

// 处理来自客户端的消息
func handleClientMessages(name string, conn *websocket.Conn) error {
	_, msg, err := conn.ReadMessage()
	if err != nil {
		return err
	}
	user := mysql.UserInfo(name)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("时间：%s\n用户：%s\n消息内容：%v\n", timestamp, name, string(msg)) // 将 msg 转换为字符串
	//broadcastMessage(user, msgType, msg)
	var data models.Data
	err = json.Unmarshal([]byte(msg), &data)
	if err != nil {
		fmt.Println("解析 JSON 错误123：", err)
	}
	if data.MsgType == "addFriends" {
		fmt.Printf("用户名:%s,添加好友名:%s\n", data.Username, data.Friendname)
		if i, err := mysql.AddFriends(data); err != nil {
			sendResponse(conn, data.MsgType, 7, "没有此人！！！")
		} else if i == 1 {
			sendResponse(conn, data.MsgType, 0, "重复申请或已经添加！！！")
		} else {
			data.Message = "5"
			SendMessageAddUser(data, timestamp) // 更新被申请人列表
			sendResponse(conn, data.MsgType, 0, "申请成功！！！")
		}
	} else if data.MsgType == "userSay" {
		// 将消息传到群里 userSay
		BroadcastMessage(user, data.Message, timestamp)
	} else if data.MsgType == "agree" {
		data.Message = "6"
		SendMessageAddUser(data, timestamp) // 更新被申请人列表
	} else if data.MsgType == "say" {
		fmt.Println(data.User, data.FriendInfo, data.Message)
		SendMessageToUser(data, timestamp)
	}
	return nil
}

// 发送消息给特定用户
func SendMessageAddUser(data models.Data, timestamp string) {
	mu.Lock()
	defer mu.Unlock()

	chatMessage := models.UFMessage{
		User:       data.User,
		FriendInfo: data.FriendInfo,
		Message:    data.Message,
		Timestamp:  timestamp,
		MsgType:    data.MsgType,
	}
	messageBytes, err := json.Marshal(chatMessage)
	if err != nil {
		fmt.Printf("消息序列化失败: %v\n", err)
		return
	}
	// 发送给目标用户
	if targetConn, ok := clients[data.Friendname]; ok {
		if err := targetConn.WriteMessage(websocket.TextMessage, messageBytes); err != nil {
			fmt.Println("发送消息失败:", err)
			return
		}
	} else {
		fmt.Printf("用户 %s 不存在\n", data.Friendname)
		return
	}
}

// 发送消息给特定用户
func SendMessageToUser(data models.Data, timestamp string) {
	mu.Lock()
	defer mu.Unlock()

	chatMessage := models.UFMessage{
		User:       data.User,
		FriendInfo: data.FriendInfo,
		Message:    data.Message,
		Timestamp:  timestamp,
		MsgType:    data.MsgType,
	}
	messageBytes, err := json.Marshal(chatMessage)
	if err != nil {
		fmt.Printf("消息序列化失败: %v\n", err)
		return
	}
	// 发送给目标用户
	if targetConn, ok := clients[data.FriendInfo.FriendName]; ok {
		if err := targetConn.WriteMessage(websocket.TextMessage, messageBytes); err != nil {
			fmt.Println("发送消息失败:", err)
			return
		}
	} else {
		fmt.Printf("用户 %s 不存在\n", data.FriendInfo.FriendName)
		return
	}

	chatMessage.MsgType = "sayUser"
	messageBytes2, err := json.Marshal(chatMessage)
	if err != nil {
		fmt.Printf("消息序列化失败: %v\n", err)
		return
	}
	fmt.Println()
	// 发送给自己
	if senderConn, ok := clients[data.User.Username]; ok {
		if err := senderConn.WriteMessage(websocket.TextMessage, messageBytes2); err != nil {
			fmt.Println("发送消息给自己失败:", err)
			return
		}
	} else {
		fmt.Printf("发送者 %s 不存在\n", data.User.Username)
		return
	}
}

// 广播消息给所有用户
func BroadcastMessage(user models.User, msg string, timestamp string) {
	mu.Lock()
	defer mu.Unlock()

	chatMessage := models.ChatMessage{
		Username:  user.Username,
		Message:   msg,
		AvatarURL: user.AvatarURL,
		Timestamp: timestamp,
	}

	if user.Username == "" {
		chatMessage.Username = "系统"
		chatMessage.AvatarURL = "10.jpg"
	}
	messageBytes, err := json.Marshal(chatMessage)
	if err != nil {
		fmt.Printf("消息序列化失败: %v\n", err)
		return
	}

	for username, conn := range clients {
		if err := conn.WriteMessage(websocket.TextMessage, messageBytes); err != nil {
			fmt.Printf("用户 %s 发送消息失败: %v\n", username, err)
		}
	}
}

func sendResponse(conn *websocket.Conn, messageID string, code int, content string) error {
	response := models.WSMessage{
		ID:      messageID,
		Code:    code,
		Content: content,
	}
	return conn.WriteJSON(response) // 写入 JSON 消息
}

// 发送在线用户列表给所有用户
// func sendOnlineUsers() {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	var onlineUsers []string
// 	for user := range clients {
// 		onlineUsers = append(onlineUsers, user)
// 	}

// 	// 将在线用户列表转换为 JSON 格式
// 	msg, err := json.Marshal(onlineUsers)
// 	if err != nil {
// 		fmt.Println("JSON 编码失败:", err)
// 		return
// 	}

// 	// 向所有用户发送在线用户列表
// 	for _, conn := range clients {
// 		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
// 			fmt.Println("发送在线用户列表失败:", err)
// 		}
// 	}
// }
