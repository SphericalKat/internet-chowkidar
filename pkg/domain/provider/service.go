// Code generated by tools/update-providers.go; DO NOT EDIT.
package di

import (
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/service"
	"github.com/google/wire"
)

type Services struct {
	service.BlocksService
	service.CategoriesService
	service.IspService
	service.SitesService
}

type MockServices struct {
}

var ServiceSet = wire.NewSet(
	service.ProvideBlocksService,
	service.ProvideCategoriesService,
	service.ProvideIspService,
	service.ProvideSitesService,
	wire.Struct(new(Services), "*"),
)

var MockServicesInstance MockServices = MockServices{}

var MockServiceSet = wire.NewSet(
	service.ProvideBlocksService,
	service.ProvideCategoriesService,
	service.ProvideIspService,
	service.ProvideSitesService,
	wire.Struct(new(Services), "*"),
	wire.Struct(new(MockServices), "*"),
)
