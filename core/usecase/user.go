package usecase

import (
	"context"
	"websocket_p4/common/log"
	"websocket_p4/common/utils"
	"websocket_p4/core/adapter/domain"
	"websocket_p4/core/entities"
)

type UseCaseUse struct {
	user domain.RepositoryUser
}

func NewUseCaseUser(
	user domain.RepositoryUser,
) *UseCaseUse {
	return &UseCaseUse{
		user: user,
	}
}

func (i *UseCaseUse) AddAccount(ctx context.Context, req *entities.User) (*entities.UserRespAdd, error) {

	user, err := i.user.FindByUserName(ctx, req.UserName)
	log.Infof("resp : ", user)
	if err != nil {
		return &entities.UserRespAdd{
			Result: entities.Result{
				Code:    1,
				Message: "error db",
			},
		}, err
	}
	if user != nil {
		return &entities.UserRespAdd{
			Result: entities.Result{
				Code:    5,
				Message: "tai khoan da ton tai vui long tao lai",
			},
		}, err
	}

	resp := utils.SetByCurlImage(ctx, req.File)
	if resp.Result.Code == 0 {
		err := i.user.AddAccount(ctx, &domain.Users{
			ID:        utils.GenerateUniqueUUid(),
			UserName:  req.UserName,
			Age:       req.Age,
			Address:   req.Address,
			Email:     req.Email,
			Avatar:    resp.URL,
			CreatedAt: utils.GetCurrentTimestamp(),
		})
		if err != nil {
			return &entities.UserRespAdd{
				Result: entities.Result{
					Code:    1,
					Message: "error db",
				},
			}, err
		}

		return &entities.UserRespAdd{
			Result: entities.Result{
				Code:    0,
				Message: "sucess",
			},
			Url:       resp.URL,
			CreatedAt: utils.GenerateUniqueUUid(),
		}, nil
	}
	return &entities.UserRespAdd{
		Result: entities.Result{
			Code:    4,
			Message: "Lỗi server",
		},
	}, nil

}
func (i *UseCaseUse) FindByUserName(ctx context.Context, user_name string) (*entities.UserRespFindByUserName, error) {

	user, err := i.user.FindByUserName(ctx, user_name)
	if err != nil {
		return &entities.UserRespFindByUserName{
			Result: entities.Result{
				Code:    1,
				Message: "error db",
			},
		}, err
	}
	if user == nil {
		return &entities.UserRespFindByUserName{
			Result: entities.Result{
				Code:    0,
				Message: "not found user",
			},
		}, err
	}
	return &entities.UserRespFindByUserName{
		Result: entities.Result{
			Code:    0,
			Message: "ok",
		},
		CreatedAt: utils.GenerateUniqueUUid(),
		User:      user,
	}, nil
}
