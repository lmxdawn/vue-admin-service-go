package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrNotFound is user not found.
	ErrNotFound = errors.NotFound("USER_NOT_FOUND", "user not found")
)

// Role is a Role model.
type Role struct {
	Name string
}

// RoleRepo is a Role repo.
type RoleRepo interface {
	Save(context.Context, *Role) (*Role, error)
	Update(context.Context, *Role) (*Role, error)
	FindByID(context.Context, int64) (*Role, error)
	ListByHello(context.Context, string) ([]*Role, error)
	ListAll(context.Context) ([]*Role, error)
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
	return uc.repo.Save(ctx, g)
}
