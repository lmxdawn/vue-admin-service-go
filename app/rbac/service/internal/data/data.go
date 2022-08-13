package data

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"vue-admin/app/rbac/service/internal/conf"
	"vue-admin/app/rbac/service/internal/data/ent"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAdminRepo, NewRoleRepo)

// Data .
type Data struct {
	db  *ent.Client
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	helper := log.NewHelper(logger)

	drv, err := sql.Open(
		c.Database.Driver,
		c.Database.Source,
	)

	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		helper.WithContext(ctx).Info(i...)
	})
	client := ent.NewClient(ent.Driver(sqlDrv))
	if err != nil {
		helper.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	// 自动创建表结构
	if err := client.Schema.Create(context.Background()); err != nil {
		helper.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	d := &Data{
		db:  client,
		rdb: rdb,
	}
	return d, func() {
		helper.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			helper.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			helper.Error(err)
		}
	}, nil
}
