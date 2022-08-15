package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
//ErrUserNotFound = errors.NotFound("USER_NOT_FOUND", "user not found")
)

// RoleAdmin is a RoleAdmin model.
type RoleAdmin struct {
	RoleId  int
	AdminId int64
}

// RoleAdminRepo is a RoleAdmin repo.
type RoleAdminRepo interface {
	SaveAll(ctx context.Context, data []*RoleAdmin) ([]int64, error)
	DeleteByAdminId(ctx context.Context, adminId int64) (int, error)
	ListByRoleId(ctx context.Context, roleId int) ([]*RoleAdmin, error)
	ListByAdminIdIn(ctx context.Context, adminIds []int64) ([]*RoleAdmin, error)
}

// RoleAdminUsecase is RoleAdmin usecase.
type RoleAdminUsecase struct {
	repo RoleAdminRepo
	log  *log.Helper
}

// NewRoleAdminUsecase new RoleAdmin usecase.
func NewRoleAdminUsecase(repo RoleAdminRepo, logger log.Logger) *RoleAdminUsecase {
	return &RoleAdminUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateRoleAdmin creates RoleAdmin, and returns the new RoleAdmin.
func (uc *RoleAdminUsecase) CreateRoleAdmin(ctx context.Context, g *RoleAdmin) (*RoleAdmin, error) {
	//return uc.repo.SaveAll(ctx, g)
	return nil, nil
}
