package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"time"
	v1 "vue-admin/api/rbac/service/v1"
	"vue-admin/app/rbac/service/internal/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound("USER_NOT_FOUND", "user not found")
)

// Login is a
type Login struct {
	Id int64
}

// LoginRepo is a Login repo.
type LoginRepo interface {
	ListRuleByAdminIdCache(ctx context.Context, id int64) ([]string, error)
	SetRuleByAdminIdCache(ctx context.Context, id int64, strings []string) (int64, error)
}

// LoginUsecase is
type LoginUsecase struct {
	repo   LoginRepo
	aRepo  AdminRepo
	raRepo RoleAdminRepo
	pRepo  PermissionRepo
	prRepo PermissionRuleRepo
	log    *log.Helper
}

// NewLoginUsecase new Login usecase.
func NewLoginUsecase(
	repo LoginRepo,
	aRepo AdminRepo,
	raRepo RoleAdminRepo,
	pRepo PermissionRepo,
	prRepo PermissionRuleRepo,
	logger log.Logger,
) *LoginUsecase {
	return &LoginUsecase{
		repo:   repo,
		aRepo:  aRepo,
		raRepo: raRepo,
		pRepo:  pRepo,
		prRepo: prRepo,
		log:    log.NewHelper(logger),
	}
}

// Login creates.
func (c *LoginUsecase) Login(ctx context.Context, req *v1.LoginRequest) (*Login, error) {

	admin, err := c.aRepo.FindByUserName(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	// 匹配密码
	b := util.CheckPasswordHash(req.Pwd, admin.Password)
	if !b {
		return nil, ErrUserNotFound
	}

	adminId := admin.Id

	// 更新登录状态
	adminUp := &Admin{
		Id:            adminId,
		LastLoginIp:   req.Ip,
		LastLoginTime: time.Now(),
		UpdateTime:    time.Now(),
	}

	_, err = c.aRepo.Update(ctx, adminUp)
	if err != nil {
		return nil, err
	}

	// 登录成功后获取权限，这里面会设置到缓存
	_, err = c.ListRuleByAdminId(ctx, adminId)
	if err != nil {
		return nil, err
	}

	c.log.WithContext(ctx).Infof("Login: %v", req.Username)
	return &Login{
		Id: adminId,
	}, nil
}

// ListRuleByAdminId creates.
func (c *LoginUsecase) ListRuleByAdminId(ctx context.Context, adminId int64) ([]string, error) {

	// 超级管理员
	if adminId == 1 {
		return []string{
			"admin",
		}, nil
	}

	// 判断是否有缓存
	cacheRules, err := c.repo.ListRuleByAdminIdCache(ctx, adminId)
	if err != nil {
		return nil, err
	}

	if len(cacheRules) > 0 {
		return cacheRules, nil
	}

	// 获取角色ids
	ras, err := c.raRepo.ListByAdminIdIn(ctx, []int64{adminId})
	if err != nil {
		return nil, err
	}

	var roleIds []int
	for _, ra := range ras {
		roleIds = append(roleIds, ra.RoleId)
	}

	// 角色授权列表
	permissions, err := c.pRepo.ListByRoleIdIn(ctx, roleIds)
	if err != nil {
		return nil, err
	}
	var permissionRuleIds []int
	for _, permission := range permissions {
		permissionRuleIds = append(permissionRuleIds, permission.PermissionRuleId)
	}

	// 获取授权规则
	permissionRules, err := c.prRepo.ListByIdIn(ctx, permissionRuleIds)
	if err != nil {
		return nil, err
	}

	// 获取权限列表
	var rules []string
	for _, rule := range permissionRules {
		rules = append(rules, rule.Name)
	}

	// 写入缓存
	_, err = c.repo.SetRuleByAdminIdCache(ctx, adminId, rules)
	if err != nil {
		return nil, err
	}

	return rules, nil
}
