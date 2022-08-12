package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"vue-admin/app/rbac/service/internal/biz"
)

type adminRepo struct {
	data *Data
	log  *log.Helper
}

// NewAdminRepo .
func NewAdminRepo(data *Data, logger log.Logger) biz.AdminRepo {
	return &adminRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (a *adminRepo) Save(ctx context.Context, g *biz.Admin) (*biz.Admin, error) {
	return g, nil
}

func (a *adminRepo) Update(ctx context.Context, g *biz.Admin) (*biz.Admin, error) {
	return g, nil
}

func (a *adminRepo) FindByID(context.Context, int64) (*biz.Admin, error) {
	return nil, nil
}

func (a *adminRepo) ListByHello(context.Context, string) ([]*biz.Admin, error) {
	return nil, nil
}

func (a *adminRepo) ListAll(context.Context) ([]*biz.Admin, error) {
	return nil, nil
}
