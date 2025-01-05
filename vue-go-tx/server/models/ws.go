package models

type Data struct {
	User       `json:"users"`
	FriendInfo `json:"friend"`
	Username   string `json:"username"`
	Message    string `json:"message"`
	Friendname string `json:"friendname"`
	Remark     string `json:"remark"`
	MsgType    string `json:"msgtype"`
}

type ChatMessage struct {
	Username  string `json:"username"`
	Message   string `json:"message"`
	AvatarURL string `json:"avatar_url"`
	Timestamp string `json:"timestamp"`
	MsgType   string `json:"msgtype"`
}

type UFMessage struct {
	User       `json:"users"`
	FriendInfo `json:"friend"`
	Message    string `json:"message"`
	Timestamp  string `json:"timestamp"`
	MsgType    string `json:"msgtype"`
}

type WSMessage struct {
	ID      string `json:"id"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}
