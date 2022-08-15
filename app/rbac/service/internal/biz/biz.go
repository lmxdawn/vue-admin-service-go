package biz

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"vue-admin/app/rbac/service/internal/data/ent"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewLoginUsecase, NewAdminUsecase, NewPermissionUsecase, NewPermissionRuleUsecase, NewRoleUsecase, NewRoleAdminUsecase)

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %w", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
