package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/jinzhu/copier"
	"vue-admin/app/rbac/service/internal/data/ent"
	"vue-admin/app/rbac/service/internal/data/ent/role"

	"github.com/go-kratos/kratos/v2/log"
	"vue-admin/app/rbac/service/internal/biz"
)

type roleRepo struct {
	data *Data
	log  *log.Helper
}

func (r *roleRepo) Save(ctx context.Context, data *biz.Role) (int, error) {
	po, err := r.data.db.Role.Create().
		SetName(data.Name).
		SetPid(data.Pid).
		SetStatus(data.Status).
		SetRemark(data.Remark).
		SetSort(data.Sort).
		SetCreateTime(data.CreateTime).
		SetUpdateTime(data.UpdateTime).
		Save(ctx)

	if err != nil {
		return 0, err
	}

	return po.ID, nil
}

func (r *roleRepo) Update(ctx context.Context, data *biz.Role) (int, error) {
	po, err := r.data.db.Role.UpdateOneID(data.Id).
		SetName(data.Name).
		SetPid(data.Pid).
		SetStatus(data.Status).
		SetRemark(data.Remark).
		SetSort(data.Sort).
		SetCreateTime(data.CreateTime).
		SetUpdateTime(data.UpdateTime).
		Save(ctx)

	if err != nil {
		return 0, err
	}

	return po.ID, nil
}

func (r *roleRepo) DeleteById(ctx context.Context, id int) error {
	return r.data.db.Role.DeleteOneID(id).Exec(ctx)
}

func (r *roleRepo) FindByName(ctx context.Context, name string) (*biz.Role, error) {
	po, err := r.data.db.Role.Query().Where(role.NameEQ(name)).Only(ctx)
	if err != nil {
		return nil, err
	}
	var res *biz.Role
	_ = copier.Copy(res, po)
	return res, nil
}

func (r *roleRepo) ListRolePage(ctx context.Context, q *biz.RoleQuery) (int, []*biz.Role, error) {
	p := r.data.db.Role.Query()
	if q.Status >= 0 {
		p.Where(role.StatusEQ(q.Status))
	}
	count, err := p.Count(ctx)
	if err != nil {
		return 0, nil, err
	}
	offset := (q.Page - 1) * q.Limit
	pos, err := p.Select(role.FieldID, role.FieldName).Offset(offset).Limit(q.Limit).Order(ent.Desc(role.FieldID)).All(ctx)
	if err != nil {
		return 0, nil, err
	}
	rv := make([]*biz.Role, 0)
	_ = copier.Copy(rv, pos)
	return count, rv, nil
}

func (r *roleRepo) ListPage(ctx context.Context, q *biz.RoleQuery) (int, []*biz.Role, error) {
	p := r.data.db.Role.Query()
	if q.Status >= 0 {
		p.Where(role.StatusEQ(q.Status))
	}
	if len(q.Name) > 0 {
		p.Where(func(selector *sql.Selector) {
			selector.Where(sql.Like(selector.C(role.FieldName), fmt.Sprintf("%%%v%%", q.Name)))
		})
	}
	count, err := p.Count(ctx)
	if err != nil {
		return 0, nil, err
	}
	offset := (q.Page - 1) * q.Limit
	pos, err := p.Offset(offset).Limit(q.Limit).Order(ent.Desc(role.FieldID)).All(ctx)
	if err != nil {
		return 0, nil, err
	}
	rv := make([]*biz.Role, 0)
	_ = copier.Copy(rv, pos)
	return count, rv, nil
}

// NewRoleRepo .
func NewRoleRepo(data *Data, logger log.Logger) biz.RoleRepo {
	return &roleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
