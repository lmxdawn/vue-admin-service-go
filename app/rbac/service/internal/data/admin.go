package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"vue-admin/app/rbac/service/internal/biz"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewAdminRepo .
func NewAdminRepo(data *Data, logger log.Logger) biz.AdminRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Admin) (*biz.Admin, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Admin) (*biz.Admin, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Admin, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Admin, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Admin, error) {
	return nil, nil
}
