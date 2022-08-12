package service

import (
	"context"
	v1 "vue-admin/api/ui/admin/v1"
	"vue-admin/app/ui/admin/internal/biz"
)

// CreateAdmin implements
func (s *UiAdmin) CreateAdmin(ctx context.Context, in *v1.CreateAdminRequest) (*v1.CreateAdminReply, error) {
	g, err := s.rc.CreateAddress(ctx, &biz.Admin{
		Username: in.Name,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateAdminReply{Name: "Hello " + g.Username}, nil
}
