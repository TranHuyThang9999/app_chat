package domain

import "context"

type Users struct {
	ID        int64  `json:"id"`
	UserName  string `json:"user_name"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

type RepositoryUser interface {
	AddAccount(ctx context.Context, req *Users) error
	FindByUserName(ctx context.Context, user_name string) (*Users, error)
}
