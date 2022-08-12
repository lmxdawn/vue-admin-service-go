package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound("USER_NOT_FOUND", "user not found")
)

// Admin is a Admin model.
type Admin struct {
	Username string
}

// AdminRepo is a Admin repo.
type AdminRepo interface {
	Save(context.Context, *Admin) (*Admin, error)
	Update(context.Context, *Admin) (*Admin, error)
	FindByID(context.Context, int64) (*Admin, error)
	ListByHello(context.Context, string) ([]*Admin, error)
	ListAll(context.Context) ([]*Admin, error)
}

// AdminUsecase is Admin usecase.
type AdminUsecase struct {
	repo AdminRepo
	log  *log.Helper
}

// NewAdminUsecase new Admin usecase.
func NewAdminUsecase(repo AdminRepo, logger log.Logger) *AdminUsecase {
	return &AdminUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateAdmin creates Admin, and returns the new Admin.
func (uc *AdminUsecase) CreateAdmin(ctx context.Context, g *Admin) (*Admin, error) {
	uc.log.WithContext(ctx).Infof("CreateAdmin: %v", g.Username)
	return uc.repo.Save(ctx, g)
}
