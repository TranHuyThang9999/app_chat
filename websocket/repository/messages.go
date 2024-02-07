package repository

import (
	"context"
	"websocket_p4/common/configs"
	"websocket_p4/core/infrastructure"
	"websocket_p4/websocket/adapter"

	"gorm.io/gorm"
)

type CollectionChat struct {
	collection *gorm.DB
}

// AddUserToChat implements adapter.RepositoryChat.
func (*CollectionChat) AddUserToChat(user adapter.Users) error {
	panic("unimplemented")
}

// GetChatHistory implements adapter.RepositoryChat.
func (*CollectionChat) GetChatHistory(user_name string) ([]*adapter.ChatMessages, error) {
	panic("unimplemented")
}

// GetMessages implements adapter.RepositoryChat.
func (*CollectionChat) GetMessages(ctx context.Context) ([]*adapter.ChatMessages, error) {
	panic("unimplemented")
}

// RemoveUserFromChat implements adapter.RepositoryChat.
func (*CollectionChat) RemoveUserFromChat(user adapter.Users) error {
	panic("unimplemented")
}

// SendUser implements adapter.RepositoryChat.
func (*CollectionChat) SendUser(ctx context.Context) error {
	panic("unimplemented")
}

func NewChatRepository(cf *configs.Configs, emp *infrastructure.PostGresql) adapter.RepositoryChat {
	return &CollectionChat{
		collection: emp.CreateCollection(),
	}
}
