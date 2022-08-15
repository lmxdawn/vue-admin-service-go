package service

import (
	"context"
	"github.com/jinzhu/copier"
	v1 "vue-admin/api/rbac/service/v1"
	"vue-admin/app/rbac/service/internal/biz"
)

// AdminService is a greeter service.
type AdminService struct {
	v1.UnimplementedAdminServer

	uc *biz.AdminUsecase
}

// NewAdminService new a service.
func NewAdminService(uc *biz.AdminUsecase) *AdminService {
	return &AdminService{uc: uc}
}

// ListAdmin implements
func (s *AdminService) ListAdmin(ctx context.Context, in *v1.ListAdminRequest) (*v1.ListAdminReply, error) {

	var req *biz.AdminQuery
	_ = copier.Copy(req, in)

	total, list, err := s.uc.List(ctx, req)
	if err != nil {
		return nil, err
	}

	var replyAdmins []*v1.ListAdminReply_Admin

	_ = copier.Copy(replyAdmins, list)

	return &v1.ListAdminReply{
		Total: int32(total),
		List:  replyAdmins,
	}, nil
}

// CreateAdmin implements
func (s *AdminService) CreateAdmin(ctx context.Context, in *v1.CreateAdminRequest) (*v1.CreateAdminReply, error) {
	g, err := s.uc.Create(ctx, &biz.Admin{Username: in.Username})
	if err != nil {
		return nil, err
	}

	var res *v1.CreateAdminReply

	_ = copier.Copy(res, g)

	return res, nil
}
