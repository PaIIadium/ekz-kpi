// +build wireinject

package main

import (
	"github.com/google/wire"
)

func ComposeApiServer(port HttpPortNumber) (*apiServer, error) {
	wire.Build(
		providers,
		wire.Struct(new(apiServer), "port", "handler"),
	)
	return nil, nil
}
