package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

var (
// ErrUserNotFound is user not found.
//ErrUserNotFound = errors.NotFound("USER_NOT_FOUND", "user not found")
)

// Admin is a Admin model.
type Admin struct {
	Id            int64
	Username      string
	Password      string
	Tel           string
	Email         string
	Avatar        string
	Sex           int
	LastLoginIp   string
	LastLoginTime time.Time
	Status        int
	CreateTime    time.Time
	UpdateTime    time.Time
}

// AdminQuery is a Admin model.
type AdminQuery struct {
	Page     int
	Limit    int
	Username string
	Status   int
	Ids      []int64
}

// AdminRepo is a Admin repo.
type AdminRepo interface {
	Save(ctx context.Context, admin *Admin) (int64, error)
	Update(ctx context.Context, admin *Admin) (int64, error)
	DeleteById(ctx context.Context, id int64) error
	FindByUserName(ctx context.Context, userName string) (*Admin, error)
	FindPwdById(ctx context.Context, id int64) (*Admin, error)
	FindById(ctx context.Context, id int64) (*Admin, error)
	ListPage(ctx context.Context, aq *AdminQuery) (int, []*Admin, error)
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

// List creates Admin, and returns the new Admin.
func (c *AdminUsecase) List(ctx context.Context, q *AdminQuery) (int, []*Admin, error) {
	return c.repo.ListPage(ctx, q)
}

// Create creates Admin, and returns the new Admin.
func (c *AdminUsecase) Create(ctx context.Context, g *Admin) (int64, error) {
	c.log.WithContext(ctx).Infof("CreateAdmin: %v", g.Username)
	return c.repo.Save(ctx, g)
}
