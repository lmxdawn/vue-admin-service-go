package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"time"
)

var (
	// ErrAdminRepeat is
	ErrAdminRepeat = errors.NotFound("ADMIN_REPEAT", "管理员已存在")
)

// Admin is
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

// AdminQuery is
type AdminQuery struct {
	Page     int
	Limit    int
	Username string
	Status   int
	Ids      []int64
}

// AdminRes is
type AdminRes struct {
	Id            int64
	Username      string
	Tel           string
	Email         string
	Avatar        string
	Sex           int
	LastLoginIp   string
	LastLoginTime time.Time
	Status        int
	CreateTime    time.Time
	UpdateTime    time.Time
	Roles         []int
}

// AdminRoleRes is
type AdminRoleRes struct {
	Id   int64
	Name string
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
	repo   AdminRepo
	rARepo RoleAdminRepo
	rRepo  RoleRepo
	log    *log.Helper
}

// NewAdminUsecase new Admin usecase.
func NewAdminUsecase(repo AdminRepo, rARepo RoleAdminRepo, rRepo RoleRepo, logger log.Logger) *AdminUsecase {
	return &AdminUsecase{repo: repo, rARepo: rARepo, rRepo: rRepo, log: log.NewHelper(logger)}
}

// List creates
func (c *AdminUsecase) List(ctx context.Context, q *AdminQuery) (int, []*AdminRes, error) {

	total, list, err := c.repo.ListPage(ctx, q)
	if err != nil {
		return 0, nil, nil
	}

	var adminIds []int64
	for _, admin := range list {
		adminIds = append(adminIds, admin.Id)
	}

	roleAdmins, err := c.rARepo.ListByAdminIdIn(ctx, adminIds)
	if err != nil {
		return 0, nil, err
	}
	roleAdminsMap := make(map[int64][]int)
	for _, roleAdmin := range roleAdmins {
		mapItem, _ := roleAdminsMap[roleAdmin.AdminId]
		mapItem = append(mapItem, roleAdmin.RoleId)
		roleAdminsMap[roleAdmin.AdminId] = mapItem
	}

	var listRes []*AdminRes
	for _, admin := range list {
		temp := &AdminRes{}
		_ = copier.Copy(temp, admin)
		if mapItem, ok := roleAdminsMap[temp.Id]; ok {
			temp.Roles = mapItem
		}
		listRes = append(listRes, temp)
	}

	return total, listRes, nil
}

// RoleList creates
func (c *AdminUsecase) RoleList(ctx context.Context, page, limit int) (int, []*Role, error) {

	return c.rRepo.ListRolePage(ctx, &RoleQuery{
		Page:  page,
		Limit: limit,
	})
}

// Create creates
func (c *AdminUsecase) Create(ctx context.Context, g *Admin) (int64, error) {
	c.log.WithContext(ctx).Infof("CreateAdmin: %v", g.Username)

	// 判断是否存在
	_, err := c.repo.FindByUserName(ctx, g.Username)
	if err == nil {
		return 0, ErrAdminRepeat
	}

	// 插入角色

	return c.repo.Save(ctx, g)
}

// Update creates
func (c *AdminUsecase) Update(ctx context.Context, g *Admin) (int64, error) {

	// 判断是否存在
	admin, err := c.repo.FindByUserName(ctx, g.Username)
	if err == nil && admin.Id != g.Id {
		return 0, ErrAdminRepeat
	}

	return c.repo.Update(ctx, g)
}
