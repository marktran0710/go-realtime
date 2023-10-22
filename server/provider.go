package server

import (
	"server/internal/user"

	"github.com/google/wire"
)

func Provider() *wire.ProviderSet {
	a := wire.NewSet(user.NewHandler, user.NewService, user.NewRepository)
	return &a
}
