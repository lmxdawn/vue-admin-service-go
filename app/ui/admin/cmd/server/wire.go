//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"vue-admin/app/ui/admin/internal/biz"
	"vue-admin/app/ui/admin/internal/conf"
	"vue-admin/app/ui/admin/internal/data"
	"vue-admin/app/ui/admin/internal/server"
	"vue-admin/app/ui/admin/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Auth, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
