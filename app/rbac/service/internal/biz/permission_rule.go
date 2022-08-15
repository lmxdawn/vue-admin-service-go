package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
//ErrUserNotFound = errors.NotFound("USER_NOT_FOUND", "user not found")
)

// PermissionRule is a PermissionRule model.
type PermissionRule struct {
	Id         int
	Pid        int
	Name       string
	Title      string
	Status     int
	Condition  string
	Sort       int
	CreateTime time.Time
	UpdateTime time.Time
}

// PermissionRuleRepo is a PermissionRule repo.
type PermissionRuleRepo interface {
	Save(ctx context.Context, data *PermissionRule) (int, error)
	Update(ctx context.Context, data *PermissionRule) (int, error)
	DeleteById(ctx context.Context, id int) error
	FindByName(ctx context.Context, name string) (*PermissionRule, error)
	ListAll(ctx context.Context) ([]*PermissionRule, error)
	ListByPid(ctx context.Context, pid int) ([]*PermissionRule, error)
	ListByIdIn(ctx context.Context, ids []int) ([]*PermissionRule, error)
}

// PermissionRuleUsecase is PermissionRule usecase.
type PermissionRuleUsecase struct {
	repo PermissionRuleRepo
	log  *log.Helper
}

// NewPermissionRuleUsecase new PermissionRule usecase.
func NewPermissionRuleUsecase(repo PermissionRuleRepo, logger log.Logger) *PermissionRuleUsecase {
	return &PermissionRuleUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreatePermissionRule creates PermissionRule, and returns the new PermissionRule.
func (uc *PermissionRuleUsecase) CreatePermissionRule(ctx context.Context, g *PermissionRule) (*PermissionRule, error) {
	//return uc.repo.SaveAll(ctx, g)
	return nil, nil
}
