package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound("", "user not found")
)

type Admin struct {
	Id       int64
	Username string
}

type AdminRepo interface {
	CreateAdmin(ctx context.Context, a *Admin) (*Admin, error)
}

type AdminUseCase struct {
	repo AdminRepo
	log  *log.Helper
}

func NewAdminUseCase(repo AdminRepo, logger log.Logger) *AdminUseCase {
	log := log.NewHelper(log.With(logger, "module", "usecase/interface"))
	return &AdminUseCase{
		repo: repo,
		log:  log,
	}
}

func (ra *AdminUseCase) CreateAdmin(ctx context.Context, a *Admin) (*Admin, error) {
	return ra.repo.CreateAdmin(ctx, a)
}
