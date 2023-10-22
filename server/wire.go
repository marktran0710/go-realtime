//go:build wireinject

// wire.go

package server

import (
	"database/sql"
	"server/internal/user"

	"github.com/google/wire"
)

func InitializeHandler(*sql.DB) (*user.Handler, error) {
	wire.Build(user.ProviderRepository, user.NewRepository)
	return &user.Handler{}, nil
}

// var Set = wire.NewSet(
// 	wire.Build(new(user.Repository), new(*Repository)))
