package loader

import (
	"websocket_p4/api/controller"
	"websocket_p4/api/routers"
	"websocket_p4/core/adapter"
	"websocket_p4/core/adapter/repository"
	"websocket_p4/core/usecase"

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
		fx.Provide(adapter.NewpostgreDb),
		fx.Provide(repository.NewEmployeesRepository),
	}
}
