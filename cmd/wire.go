//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/google/wire"

	"github.com/devhg/es/internal/biz"
	"github.com/devhg/es/internal/conf"
	"github.com/devhg/es/internal/data"
	"github.com/devhg/es/internal/server"
	"github.com/devhg/es/internal/service"
)

// The build tag makes sure the stub is not built in the final build.

// initApp init application.
func initApp(*conf.Config) (*http.Server, func()) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet))
}
