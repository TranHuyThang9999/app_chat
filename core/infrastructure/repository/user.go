package repository

import (
	"context"
	"errors"
	"websocket_p4/common/configs"
	"websocket_p4/core/infrastructure"
	"websocket_p4/core/infrastructure/domain"

	"gorm.io/gorm"
)

type CollectionUser struct {
	collection *gorm.DB
}

func NewEmployeesRepository(cf *configs.Configs, emp *infrastructure.PostGresql) domain.RepositoryUser {
	return &CollectionUser{
		collection: emp.CreateCollection(),
	}
}

// AddAccount implements domain.RepositoryUser.
func (u *CollectionUser) AddAccount(ctx context.Context, req *domain.Users) error {
	result := u.collection.Create(&req)
	return result.Error
}

// FindByUserName implements domain.RepositoryUser.
func (u *CollectionUser) FindByUserName(ctx context.Context, user_name string) (*domain.Users, error) {
	var user *domain.Users
	result := u.collection.Where("user_name = ?", user_name).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return user, result.Error
}
