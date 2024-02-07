package loader

import (
	"websocket_p4/api/controller"
	"websocket_p4/api/routers"
	"websocket_p4/core/infrastructure"
	"websocket_p4/core/infrastructure/repository"
	"websocket_p4/core/usecase"
	"websocket_p4/websocket"

	"go.uber.org/fx"
)

func Load() []fx.Option {
	return []fx.Option{
		fx.Options(loadAdapter()...),
		fx.Options(loadUseCase()...),
		fx.Options(loadValidator()...),
		fx.Options(loadEngine()...),
	}
}
func loadUseCase() []fx.Option {
	return []fx.Option{
		fx.Provide(usecase.NewUseCaseUser),
		fx.Provide(websocket.NewServer),
		fx.Provide(websocket.NewManagerClient),
	}
}

func loadValidator() []fx.Option {
	return []fx.Option{
		//	fx.Provide(validator.New),
	}
}
func loadEngine() []fx.Option {
	return []fx.Option{
		fx.Provide(routers.NewApiRouter),
		// fx.Provide(controllers.NewBaseController),
		fx.Provide(controller.NewControllerEmployees),
	}
}
func loadAdapter() []fx.Option {
	return []fx.Option{
		fx.Provide(infrastructure.NewpostgreDb),
		fx.Provide(repository.NewEmployeesRepository),
	}
}
