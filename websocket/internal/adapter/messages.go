package adapter

import (
	"context"
	"websocket_p4/common/configs"
	"websocket_p4/core/infrastructure"
	"websocket_p4/websocket/internal/repository"

	"gorm.io/gorm"
)

// ChatMessage đại diện cho một tin nhắn trong chat

type CollectionChat struct {
	collection *gorm.DB
}

// AddUserToChat implements repository.RepositoryChat.
func (*CollectionChat) AddUserToChat(user repository.Users) error {
	panic("unimplemented")
}

// GetChatHistory implements repository.RepositoryChat.
func (*CollectionChat) GetChatHistory(user_name string) ([]*repository.ChatMessages, error) {
	panic("unimplemented")
}

// GetMessages implements repository.RepositoryChat.
func (*CollectionChat) GetMessages(ctx context.Context) ([]*repository.ChatMessages, error) {
	panic("unimplemented")
}

// RemoveUserFromChat implements repository.RepositoryChat.
func (*CollectionChat) RemoveUserFromChat(user repository.Users) error {
	panic("unimplemented")
}

// SendUser implements repository.RepositoryChat.
func (*CollectionChat) SendUser(ctx context.Context) error {
	panic("unimplemented")
}

func NewChatRepository(cf *configs.Configs, emp *infrastructure.PostGresql) repository.RepositoryChat {
	return &CollectionChat{
		collection: emp.CreateCollection(),
	}
}
