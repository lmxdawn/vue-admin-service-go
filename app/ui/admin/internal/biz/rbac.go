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

type RbacRepo interface {
	CreateAdmin(ctx context.Context, admin *Admin) (*Admin, error)
}

type RbacUseCase struct {
	repo RbacRepo
	log  *log.Helper
}

func NewRbacUseCase(repo RbacRepo, logger log.Logger) *RbacUseCase {
	log := log.NewHelper(log.With(logger, "module", "usecase/interface"))
	return &RbacUseCase{
		repo: repo,
		log:  log,
	}
}

func (rc *RbacUseCase) CreateAddress(ctx context.Context, admin *Admin) (*Admin, error) {
	return rc.repo.CreateAdmin(ctx, admin)
}
