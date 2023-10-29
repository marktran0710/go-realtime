//go:build wireinject

// wire.go

package server

import (
	"server/internal/user"

	"github.com/google/wire"
)

func InitializeHandler(db user.DBTX) (*user.Handler, error) {
	wire.Build(user.NewHandler, user.NewRepository, user.NewService)
	return &user.Handler{}, nil
}
