package usecase

import (
	"context"
	"websocket_p4/common/log"
	"websocket_p4/common/utils"
	"websocket_p4/core/infrastructure/domain"
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
		}, nil
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
			Code:    resp.Result.Code,
			Message: resp.Result.Message,
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
				Code:    7,
				Message: "not found user",
			},
		}, nil
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
func (i *UseCaseUse) Login(ctx context.Context, req *entities.UserReqLogin) (*entities.UserRespLogin, error) {

	//log.Infof("req : ", req.UserName)

	user, err := i.user.FindByUserName(ctx, req.UserName)
	if err != nil {
		return &entities.UserRespLogin{
			Result: entities.Result{
				Code:    1,
				Message: "error db",
			},
		}, nil
	}
	if user == nil {
		return &entities.UserRespLogin{
			Result: entities.Result{
				Code:    7,
				Message: "not found user",
			},
		}, nil
	}

	if user.Email != req.Email {
		return &entities.UserRespLogin{
			Result: entities.Result{
				Code:    7,
				Message: "not found user",
			},
		}, nil
	}
	//	log.Infof("user :", user)
	return &entities.UserRespLogin{
		Result: entities.Result{
			Code:    0,
			Message: "login sucess",
		},
	}, nil
}
