package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"vue-admin/app/rbac/service/internal/biz"
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

func (a *roleRepo) Save(ctx context.Context, g *biz.Role) (*biz.Role, error) {
	return g, nil
}

func (a *roleRepo) Update(ctx context.Context, g *biz.Role) (*biz.Role, error) {
	return g, nil
}

func (a *roleRepo) FindByID(context.Context, int64) (*biz.Role, error) {
	return nil, nil
}

func (a *roleRepo) ListByHello(context.Context, string) ([]*biz.Role, error) {
	return nil, nil
}

func (a *roleRepo) ListAll(context.Context) ([]*biz.Role, error) {
	return nil, nil
}
