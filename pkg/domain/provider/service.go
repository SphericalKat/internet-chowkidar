// Code generated by tools/update-providers.go; DO NOT EDIT.
package di

import (
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/service"
	mock_service "github.com/gnulinuxindia/internet-chowkidar/pkg/domain/service/mock"
	"github.com/google/wire"
	"go.uber.org/mock/gomock"
)

type Services struct {
	service.CounterService
	service.EmailService
}

type MockServices struct {
	*mock_service.MockEmailService
}

var ServiceSet = wire.NewSet(
	service.ProvideCounterService,
	service.ProvideEmailService,
	wire.Struct(new(Services), "*"),
)

func ProvideMockEmailService(ctrl *gomock.Controller) service.EmailService {
	MockServicesInstance.MockEmailService = mock_service.NewMockEmailService(ctrl)
	return MockServicesInstance.MockEmailService
}

var MockServicesInstance MockServices = MockServices{}

var MockServiceSet = wire.NewSet(
	service.ProvideCounterService,
	ProvideMockEmailService,
	wire.Struct(new(Services), "*"),
	wire.Struct(new(MockServices), "*"),
)
