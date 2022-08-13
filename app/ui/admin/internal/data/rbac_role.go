package data

import (
	"context"
	rbacv1 "vue-admin/api/rbac/service/v1"

	"github.com/go-kratos/kratos/v2/log"
	"vue-admin/app/ui/admin/internal/biz"
)

type roleRepo struct {
	data *Data
	log  *log.Helper
}

// NewRoleRepo .
func NewRoleRepo(data *Data, logger log.Logger) biz.RoleRepo {
	return &roleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *roleRepo) CreateRole(ctx context.Context, role *biz.Role) (*biz.Role, error) {

	res, err := r.data.rr.CreateRole(ctx, &rbacv1.CreateRoleRequest{
		Name: role.Name,
	})

	if err != nil {
		return nil, err
	}

	return &biz.Role{
		Id:   1,
		Name: res.Name,
	}, nil
}
