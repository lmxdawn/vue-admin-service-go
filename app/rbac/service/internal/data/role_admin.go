package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"vue-admin/app/rbac/service/internal/biz"
	"vue-admin/app/rbac/service/internal/data/ent"
	"vue-admin/app/rbac/service/internal/data/ent/roleadmin"
)

type roleAdminRepo struct {
	data *Data
	log  *log.Helper
}

func (r *roleAdminRepo) SaveAll(ctx context.Context, data []*biz.RoleAdmin) ([]int64, error) {
	bulk := make([]*ent.RoleAdminCreate, len(data))
	for i, item := range data {
		bulk[i] = r.data.db.RoleAdmin.Create().
			SetRoleID(item.RoleId).
			SetAdminID(item.AdminId)
	}
	po, err := r.data.db.RoleAdmin.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, err
	}
	rs := make([]int64, len(po))
	for i, p := range po {
		rs[i] = p.ID
	}
	return rs, nil
}

func (r *roleAdminRepo) DeleteByAdminId(ctx context.Context, adminId int64) (int, error) {
	return r.data.db.RoleAdmin.Delete().Where(roleadmin.AdminIDEQ(adminId)).Exec(ctx)
}

func (r *roleAdminRepo) ListByRoleId(ctx context.Context, roleId int) ([]*biz.RoleAdmin, error) {
	po, err := r.data.db.RoleAdmin.Query().Where(roleadmin.RoleIDEQ(roleId)).All(ctx)
	if err != nil {
		return nil, err
	}
	var res []*biz.RoleAdmin
	_ = copier.Copy(res, po)
	return res, nil
}

func (r *roleAdminRepo) ListByAdminIdIn(ctx context.Context, adminIds []int64) ([]*biz.RoleAdmin, error) {
	po, err := r.data.db.RoleAdmin.Query().Where(roleadmin.AdminIDIn(adminIds...)).All(ctx)
	if err != nil {
		return nil, err
	}
	var res []*biz.RoleAdmin
	_ = copier.Copy(res, po)
	return res, nil
}

// NewRoleAdminRepo .
func NewRoleAdminRepo(data *Data, logger log.Logger) biz.RoleAdminRepo {
	return &roleAdminRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
