package adapter

import "context"

// ChatMessage đại diện cho một tin nhắn trong chat
type ChatMessages struct {
	Id        int64  `json:"id"`
	IdUser    int64  `json:"id_user"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// User đại diện cho một người dùng trong chat

type Users struct {
	Id        int64  `json:"id"`
	Avatar    string `json:"avatar"`
	Status    int    `json:"status"`
	CreatedAt int64  `json:"created_at"`
}

type RepositoryChat interface {
	SendUser(ctx context.Context) error
	GetMessages(ctx context.Context) ([]*ChatMessages, error)
	GetChatHistory(user_name string) ([]*ChatMessages, error)
	AddUserToChat(user Users) error
	RemoveUserFromChat(user Users) error
}
