//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/skyrocketOoO/serverx/internal/controller"
	authC "github.com/skyrocketOoO/serverx/internal/controller/auth"
	generalC "github.com/skyrocketOoO/serverx/internal/controller/general"
	"github.com/skyrocketOoO/serverx/internal/service/cognito"
	authU "github.com/skyrocketOoO/serverx/internal/usecase/auth"
	generalU "github.com/skyrocketOoO/serverx/internal/usecase/general"
)

func NewHandler() (*controller.Handler, error) {
	wire.Build(
		controller.NewHandler,
		generalC.NewHandler,
		authC.NewHandler,
		wire.Bind(new(authC.Usecase), new(*authU.Usecase)),
		wire.Bind(new(generalC.Usecase), new(*generalU.Usecase)),
		authU.NewUsecase,
		generalU.NewUsecase,
		cognito.New,
	)
	return &controller.Handler{}, nil
}
