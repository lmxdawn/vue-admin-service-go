package data

import (
	"context"
	"vue-admin/app/rbac/service/internal/data/ent"
	"vue-admin/app/rbac/service/internal/data/ent/admin"
	"vue-admin/app/rbac/service/internal/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
	"vue-admin/app/rbac/service/internal/biz"

	"github.com/jinzhu/copier"
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

func (r *adminRepo) Save(ctx context.Context, admin *biz.Admin) (int64, error) {
	ph, err := util.HashPassword(admin.Password)
	if err != nil {
		return 0, err
	}
	po, err := r.data.db.Admin.
		Create().
		SetUsername(admin.Username).
		SetPassword(ph).
		SetTel(admin.Tel).
		SetEmail(admin.Email).
		SetAvatar(admin.Avatar).
		SetSex(admin.Sex).
		SetLastLoginIP(admin.LastLoginIp).
		SetLastLoginTime(admin.LastLoginTime).
		SetCreateTime(admin.CreateTime).
		SetUpdateTime(admin.UpdateTime).
		Save(ctx)
	if err != nil {
		return 0, err
	}
	return po.ID, nil
}

func (r *adminRepo) Update(ctx context.Context, admin *biz.Admin) (int64, error) {
	p := r.data.db.Admin.
		UpdateOneID(admin.Id).
		SetTel(admin.Tel).
		SetEmail(admin.Email).
		SetAvatar(admin.Avatar).
		SetSex(admin.Sex).
		SetLastLoginIP(admin.LastLoginIp).
		SetLastLoginTime(admin.LastLoginTime).
		SetCreateTime(admin.CreateTime).
		SetUpdateTime(admin.UpdateTime)
	if len(admin.Password) > 0 {
		ph, err := util.HashPassword(admin.Password)
		if err != nil {
			return 0, err
		}
		p.SetPassword(ph)
	}
	po, err := p.Save(ctx)
	if err != nil {
		return 0, err
	}
	return po.ID, nil
}

func (r *adminRepo) DeleteById(ctx context.Context, id int64) error {
	err := r.data.db.Admin.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *adminRepo) FindByUserName(ctx context.Context, userName string) (*biz.Admin, error) {
	po, err := r.data.db.Admin.
		Query().
		Where(admin.UsernameEQ(userName)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Admin{
		Id:       po.ID,
		Password: po.Password,
	}, nil
}

func (r *adminRepo) FindPwdById(ctx context.Context, id int64) (*biz.Admin, error) {
	po, err := r.data.db.Admin.
		Query().
		Where(admin.IDEQ(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Admin{
		Password: po.Password,
	}, nil
}

func (r *adminRepo) FindById(ctx context.Context, id int64) (*biz.Admin, error) {
	po, err := r.data.db.Admin.
		Query().
		Where(admin.IDEQ(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Admin{
		Id:       po.ID,
		Username: po.Username,
		Avatar:   po.Avatar,
	}, nil
}

func (r *adminRepo) ListPage(ctx context.Context, aq *biz.AdminQuery) (int, []*biz.Admin, error) {
	p := r.data.db.Admin.Query()
	if len(aq.Username) > 0 {
		p.Where(admin.UsernameEQ(aq.Username))
	}
	if aq.Status >= 0 {
		p.Where(admin.StatusEQ(aq.Status))
	}
	if aq.Ids != nil && len(aq.Ids) > 0 {
		p.Where(admin.IDIn(aq.Ids...))
	}
	count, err := p.Count(ctx)
	if err != nil {
		return 0, nil, err
	}
	offset := (aq.Page - 1) * aq.Limit
	pos, err := p.Offset(offset).Limit(aq.Limit).Order(ent.Desc(admin.FieldID)).All(ctx)
	if err != nil {
		return 0, nil, err
	}
	rv := make([]*biz.Admin, 0)
	_ = copier.Copy(rv, pos)
	return count, rv, nil
}
