package service

import (
	"context"
	v1 "vue-admin/api/rbac/service/v1"
	"vue-admin/app/rbac/service/internal/biz"
)

// CreateAdmin implements
func (s *RbacService) CreateAdmin(ctx context.Context, in *v1.CreateAdminRequest) (*v1.CreateAdminReply, error) {
	g, err := s.uc.CreateAdmin(ctx, &biz.Admin{Username: in.Username})
	if err != nil {
		return nil, err
	}
	return &v1.CreateAdminReply{
		Admin: &v1.CreateAdminReply_Admin{
			Id:       0,
			Username: g.Username,
		},
	}, nil
}
