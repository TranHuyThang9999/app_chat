// mapper.go
package mapper

import (
	"websocket_p4/core/entities"
	"websocket_p4/core/infrastructure/domain"
)

func ConvertUserEntityToDomain(userEntity *domain.Users) *entities.UserResponseGetAll {
	return &entities.UserResponseGetAll{
		UserName: userEntity.UserName,
		Age:      userEntity.Age,
		Address:  userEntity.Address,
		Email:    userEntity.Email,
		File:     userEntity.Avatar,
		// You may set CreatedAt based on your requirement.
	}
}

func ConvertUserEntitiesToDomainList(userEntities []*domain.Users) []*entities.UserResponseGetAll {
	var userList []*entities.UserResponseGetAll
	for _, userEntity := range userEntities {
		userList = append(userList, ConvertUserEntityToDomain(userEntity))
	}
	return userList
}
