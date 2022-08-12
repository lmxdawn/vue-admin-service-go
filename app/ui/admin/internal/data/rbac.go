package data

import (
	"context"
	rbacv1 "vue-admin/api/rbac/service/v1"

	"github.com/go-kratos/kratos/v2/log"
	"vue-admin/app/ui/admin/internal/biz"
)

type rbacRepo struct {
	data *Data
	log  *log.Helper
}

// NewRbacRepo .
func NewRbacRepo(data *Data, logger log.Logger) biz.RbacRepo {
	return &rbacRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *rbacRepo) CreateAdmin(ctx context.Context, admin *biz.Admin) (*biz.Admin, error) {

	res, err := r.data.rc.CreateAdmin(ctx, &rbacv1.CreateAdminRequest{
		Username: admin.Username,
	})

	if err != nil {
		return nil, err
	}

	return &biz.Admin{
		Id:       1,
		Username: res.Admin.Username,
	}, nil
}
