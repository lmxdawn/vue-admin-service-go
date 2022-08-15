package service

import (
	"context"
	v1 "vue-admin/api/rbac/service/v1"
	"vue-admin/app/rbac/service/internal/biz"
)

// LoginService is a greeter service.
type LoginService struct {
	v1.UnimplementedLoginServer

	uc *biz.LoginUsecase
}

// NewLoginService new a service.
func NewLoginService(uc *biz.LoginUsecase) *LoginService {
	return &LoginService{uc: uc}
}

// Login implements
func (s *LoginService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	loginInfo, err := s.uc.Login(ctx, in)
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{
		Id: loginInfo.Id,
	}, nil
}
