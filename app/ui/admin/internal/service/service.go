package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"vue-admin/app/ui/admin/internal/biz"

	"vue-admin/api/ui/admin/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUiAdmin)

type UiAdmin struct {
	v1.UnimplementedUiAdminServer

	log *log.Helper
	rc  *biz.RbacUseCase
}

func NewUiAdmin(
	rc *biz.RbacUseCase,
	logger log.Logger,
) *UiAdmin {
	return &UiAdmin{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		rc:  rc,
	}
}
