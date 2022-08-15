package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"vue-admin/app/rbac/service/internal/biz"
	"vue-admin/app/rbac/service/internal/data/ent/permissionrule"
)

type permissionRuleRepo struct {
	data *Data
	log  *log.Helper
}

func (r *permissionRuleRepo) Save(ctx context.Context, data *biz.PermissionRule) (int, error) {
	po, err := r.data.db.PermissionRule.Create().
		SetPid(data.Pid).
		SetName(data.Name).
		SetTitle(data.Title).
		SetStatus(data.Status).
		SetCondition(data.Condition).
		SetSort(data.Sort).
		SetCreateTime(data.CreateTime).
		SetUpdateTime(data.UpdateTime).
		Save(ctx)
	if err != nil {
		return 0, err
	}
	return po.ID, nil
}

func (r *permissionRuleRepo) Update(ctx context.Context, data *biz.PermissionRule) (int, error) {
	po, err := r.data.db.PermissionRule.UpdateOneID(data.Id).
		SetPid(data.Pid).
		SetName(data.Name).
		SetTitle(data.Title).
		SetStatus(data.Status).
		SetCondition(data.Condition).
		SetSort(data.Sort).
		SetCreateTime(data.CreateTime).
		SetUpdateTime(data.UpdateTime).
		Save(ctx)
	if err != nil {
		return 0, err
	}
	return po.ID, nil
}

func (r *permissionRuleRepo) DeleteById(ctx context.Context, id int) error {
	return r.data.db.PermissionRule.DeleteOneID(id).Exec(ctx)
}

func (r *permissionRuleRepo) FindByName(ctx context.Context, name string) (*biz.PermissionRule, error) {
	po, err := r.data.db.PermissionRule.Query().Where(permissionrule.NameEQ(name)).Only(ctx)
	if err != nil {
		return nil, err
	}
	var res biz.PermissionRule
	_ = copier.Copy(&res, po)
	return &res, nil
}

func (r *permissionRuleRepo) ListAll(ctx context.Context) ([]*biz.PermissionRule, error) {
	po, err := r.data.db.PermissionRule.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var res []*biz.PermissionRule
	_ = copier.Copy(res, po)
	return res, nil
}

func (r *permissionRuleRepo) ListByPid(ctx context.Context, pid int) ([]*biz.PermissionRule, error) {
	po, err := r.data.db.PermissionRule.Query().Where(permissionrule.PidEQ(pid)).All(ctx)
	if err != nil {
		return nil, err
	}
	var res []*biz.PermissionRule
	_ = copier.Copy(res, po)
	return res, nil
}

func (r *permissionRuleRepo) ListByIdIn(ctx context.Context, ids []int) ([]*biz.PermissionRule, error) {
	po, err := r.data.db.PermissionRule.Query().Where(permissionrule.IDIn(ids...)).All(ctx)
	if err != nil {
		return nil, err
	}
	var res []*biz.PermissionRule
	_ = copier.Copy(res, po)
	return res, nil
}

// NewPermissionRuleRepo .
func NewPermissionRuleRepo(data *Data, logger log.Logger) biz.PermissionRuleRepo {
	return &permissionRuleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
