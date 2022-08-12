package service

import (
	"github.com/google/wire"
	v1 "vue-admin/api/rbac/service/v1"
	"vue-admin/app/rbac/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAdminService)

// RbacService is a greeter service.
type RbacService struct {
	v1.UnimplementedRbacServer

	uc *biz.AdminUsecase
}

// NewAdminService new a greeter service.
func NewAdminService(uc *biz.AdminUsecase) *RbacService {
	return &RbacService{uc: uc}
}
