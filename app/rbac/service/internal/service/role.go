package service

import (
	"context"
	v1 "vue-admin/api/rbac/service/v1"
	"vue-admin/app/rbac/service/internal/biz"
)

// RoleService is a greeter service.
type RoleService struct {
	v1.UnimplementedRoleServer

	uc *biz.RoleUsecase
}

// NewRoleService new a service.
func NewRoleService(uc *biz.RoleUsecase) *RoleService {
	return &RoleService{uc: uc}
}

// CreateRole implements
func (s *RoleService) CreateRole(ctx context.Context, in *v1.CreateRoleRequest) (*v1.CreateRoleReply, error) {
	g, err := s.uc.CreateRole(ctx, &biz.Role{Name: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.CreateRoleReply{
		Name: g.Name,
	}, nil
}
