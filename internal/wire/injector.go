//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	authC "github.com/skyrocketOoO/serverx/internal/controller/auth"
	"github.com/skyrocketOoO/serverx/internal/service/cognito"
	authU "github.com/skyrocketOoO/serverx/internal/usecase/auth"
)

func NewHandler() (*authC.Handler, error) {
	wire.Build(
		authC.NewHandler,
		wire.Bind(new(authC.Usecase), new(*authU.Usecase)),
		authU.NewUsecase,
		cognito.New,
	)
	return &authC.Handler{}, nil
}
