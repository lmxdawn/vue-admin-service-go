package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"vue-admin/app/rbac/service/internal/biz"
	"vue-admin/app/rbac/service/internal/data/ent"
	"vue-admin/app/rbac/service/internal/data/ent/permission"
)

type permissionRepo struct {
	data *Data
	log  *log.Helper
}

func (r *permissionRepo) SaveAll(ctx context.Context, datas []*biz.Permission) ([]int, error) {
	bulk := make([]*ent.PermissionCreate, len(datas))
	for i, data := range datas {
		bulk[i] = r.data.db.Permission.Create().SetRoleID(data.RoleId).SetPermissionRuleID(data.PermissionRuleId).SetType(data.Type)
	}
	po, err := r.data.db.Permission.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, err
	}
	rs := make([]int, len(po))
	for i, p := range po {
		rs[i] = p.ID
	}
	return rs, nil
}

func (r *permissionRepo) DeleteByRoleId(ctx context.Context, roleId int) (int, error) {
	return r.data.db.Permission.Delete().Where(permission.RoleIDEQ(roleId)).Exec(ctx)
}

func (r *permissionRepo) ListByRoleId(ctx context.Context, roleId int) ([]*biz.Permission, error) {
	po, err := r.data.db.Permission.Query().Where(permission.RoleIDEQ(roleId)).All(ctx)
	if err != nil {
		return nil, err
	}
	var bizs []*biz.Permission
	_ = copier.Copy(bizs, po)
	return bizs, nil
}

func (r *permissionRepo) ListByRoleIdIn(ctx context.Context, roleIds []int) ([]*biz.Permission, error) {
	po, err := r.data.db.Permission.Query().Where(permission.RoleIDIn(roleIds...)).All(ctx)
	if err != nil {
		return nil, err
	}
	var bizs []*biz.Permission
	_ = copier.Copy(bizs, po)
	return bizs, nil
}

// NewPermissionRepo .
func NewPermissionRepo(data *Data, logger log.Logger) biz.PermissionRepo {
	return &permissionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
