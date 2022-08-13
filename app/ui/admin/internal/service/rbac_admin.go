package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "vue-admin/api/ui/admin/v1"
	"vue-admin/app/ui/admin/internal/biz"
)

type RbacAdmin struct {
	v1.UnimplementedRbacAdminServer

	log *log.Helper
	ra  *biz.AdminUseCase
}

func NewRbacAdmin(
	ra *biz.AdminUseCase,
	logger log.Logger,
) *RbacAdmin {
	return &RbacAdmin{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		ra:  ra,
	}
}

// CreateRbacAdmin implements
func (s *RbacAdmin) CreateRbacAdmin(ctx context.Context, in *v1.CreateRbacAdminRequest) (*v1.CreateRbacAdminReply, error) {
	g, err := s.ra.CreateAdmin(ctx, &biz.Admin{
		Username: in.Name,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateRbacAdminReply{Name: "Admin 请求 " + g.Username}, nil
}
