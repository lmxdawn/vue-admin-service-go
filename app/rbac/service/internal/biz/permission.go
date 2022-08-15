package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
//ErrUserNotFound = errors.NotFound("USER_NOT_FOUND", "user not found")
)

// Permission is a Permission model.
type Permission struct {
	RoleId           int
	PermissionRuleId int
	Type             string
}

// PermissionRepo is a Permission repo.
type PermissionRepo interface {
	SaveAll(ctx context.Context, datas []*Permission) ([]int, error)
	DeleteByRoleId(ctx context.Context, roleId int) (int, error)
	ListByRoleId(ctx context.Context, roleId int) ([]*Permission, error)
	ListByRoleIdIn(ctx context.Context, roleIds []int) ([]*Permission, error)
}

// PermissionUsecase is Permission usecase.
type PermissionUsecase struct {
	repo PermissionRepo
	log  *log.Helper
}

// NewPermissionUsecase new Permission usecase.
func NewPermissionUsecase(repo PermissionRepo, logger log.Logger) *PermissionUsecase {
	return &PermissionUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreatePermission creates Permission, and returns the new Permission.
func (uc *PermissionUsecase) CreatePermission(ctx context.Context, g *Permission) (*Permission, error) {
	//return uc.repo.SaveAll(ctx, g)
	return nil, nil
}
