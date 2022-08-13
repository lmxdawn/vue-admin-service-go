package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "vue-admin/api/ui/admin/v1"
	"vue-admin/app/ui/admin/internal/biz"
)

type RbacRole struct {
	v1.UnimplementedRbacRoleServer

	log *log.Helper
	rr  *biz.RoleUseCase
}

func NewRbacRole(
	rr *biz.RoleUseCase,
	logger log.Logger,
) *RbacRole {
	return &RbacRole{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		rr:  rr,
	}
}

// CreateRbacRole implements
func (s *RbacRole) CreateRbacRole(ctx context.Context, in *v1.CreateRbacRoleRequest) (*v1.CreateRbacRoleReply, error) {
	g, err := s.rr.CreateRole(ctx, &biz.Role{
		Name: in.Name,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateRbacRoleReply{Name: "CCCC " + g.Name}, nil
}
