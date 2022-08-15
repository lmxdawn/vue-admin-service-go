package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"vue-admin/app/rbac/service/internal/biz"
)

var rulesCacheKey = func(id int64) string {
	return fmt.Sprintf("rbac:rules:%d", id)
}

type LoginRepo struct {
	data *Data
	log  *log.Helper
}

func (r *LoginRepo) ListRuleByAdminIdCache(ctx context.Context, id int64) ([]string, error) {
	data := r.data.rdb.SMembers(ctx, rulesCacheKey(id))
	return data.Result()
}

func (r *LoginRepo) SetRuleByAdminIdCache(ctx context.Context, id int64, strings []string) (int64, error) {
	data := r.data.rdb.SAdd(ctx, rulesCacheKey(id), strings)
	return data.Result()
}

// NewLoginRepo .
func NewLoginRepo(data *Data, logger log.Logger) biz.LoginRepo {
	return &LoginRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
