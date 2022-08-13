package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrPasswordInvalid = errors.New("password invalid")
)

type Role struct {
	Id     int64
	Name   string
	Status string
}

type RoleRepo interface {
	CreateRole(ctx context.Context, a *Role) (*Role, error)
}

type RoleUseCase struct {
	repo RoleRepo
	log  *log.Helper
}

func NewRoleUseCase(repo RoleRepo, logger log.Logger) *RoleUseCase {
	log := log.NewHelper(log.With(logger, "module", "usecase/interface"))
	return &RoleUseCase{
		repo: repo,
		log:  log,
	}
}

func (rc *RoleUseCase) CreateRole(ctx context.Context, ra *Role) (*Role, error) {
	return rc.repo.CreateRole(ctx, ra)
}
