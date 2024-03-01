package models

// Struct cho bảng "users"
type Users struct {
	ID        int64  `json:"id"`
	UserName  string `json:"user_name"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	CreatedAt int    `json:"created_at"`
}

// Struct cho bảng "messages"
type Messages struct {
	ID          int64  `json:"id"`
	Content     string `json:"content"`
	SenderID    int64  `json:"sender_id"`
	RecipientID int64  `json:"recipient_id"`
	CreatedAt   int    `json:"created_at"`
}

// Struct cho bảng "message_status"
type MessageStatus struct {
	ID          int64 `json:"id"`
	MessageID   int64 `json:"message_id"`
	SenderID    int64 `json:"sender_id"`
	RecipientID int64 `json:"recipient_id"`
	IsSent      bool  `json:"is_sent"`
	IsReceived  bool  `json:"is_received"`
	CreatedAt   int   `json:"created_at"`
}
