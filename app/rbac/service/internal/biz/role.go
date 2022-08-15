package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrNotFound is user not found.
	ErrNotFound = errors.NotFound("USER_NOT_FOUND", "user not found")
)

// Role is a Role model.
type Role struct {
	Id         int
	Name       string
	Pid        int
	Status     int
	Remark     string
	Sort       int
	CreateTime time.Time
	UpdateTime time.Time
}

// RoleQuery is a
type RoleQuery struct {
	Page   int
	Limit  int
	Name   string
	Status int
}

// RoleRepo is a Role repo.
type RoleRepo interface {
	Save(ctx context.Context, data *Role) (int, error)
	Update(ctx context.Context, data *Role) (int, error)
	DeleteById(ctx context.Context, id int) error
	FindByName(ctx context.Context, name string) (*Role, error)
	ListRolePage(ctx context.Context, q *RoleQuery) (int, []*Role, error)
	ListPage(ctx context.Context, q *RoleQuery) (int, []*Role, error)
}

// RoleUsecase is Role usecase.
type RoleUsecase struct {
	repo RoleRepo
	log  *log.Helper
}

// NewRoleUsecase new Role usecase.
func NewRoleUsecase(repo RoleRepo, logger log.Logger) *RoleUsecase {
	return &RoleUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateRole creates Role, and returns the new Role.
func (uc *RoleUsecase) CreateRole(ctx context.Context, g *Role) (*Role, error) {
	uc.log.WithContext(ctx).Infof("CreateRole: %v", g.Name)
	//return uc.repo.Save(ctx, g)
	return nil, nil
}
