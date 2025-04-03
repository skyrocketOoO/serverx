//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	authC "github.com/skyrocketOoO/serverx/internal/controller/auth"
	authU "github.com/skyrocketOoO/serverx/internal/usecase/auth"
)

func NewHandler() *authC.Handler {
	wire.Build(
		authC.NewHandler,
		wire.Bind(new(authC.Usecase), new(*authU.Usecase)),
		authU.NewUsecase,
	)
	return &authC.Handler{}
}
