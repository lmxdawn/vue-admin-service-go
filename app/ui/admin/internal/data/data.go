package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	grpcClient "google.golang.org/grpc"
	"vue-admin/app/ui/admin/internal/conf"

	rbacV1 "vue-admin/api/rbac/service/v1"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDiscovery,
	NewRegistrar,
	NewRbacServiceConn,
	NewRbacServiceAdminClient,
	NewRbacServiceRoleClient,
	NewAdminRepo,
	NewRoleRepo,
)

// Data .
type Data struct {
	log *log.Helper
	ra  rbacV1.AdminClient
	rr  rbacV1.RoleClient
}

// NewData .
func NewData(
	conf *conf.Data,
	logger log.Logger,
	ra rbacV1.AdminClient,
	rr rbacV1.RoleClient,
) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{
		log: l,
		ra:  ra,
		rr:  rr,
	}, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewRbacServiceConn(ac *conf.Auth, r registry.Discovery) *grpcClient.ClientConn {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///rbac.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			jwt.Client(func(token *jwt2.Token) (interface{}, error) {
				return []byte(ac.ServiceKey), nil
			}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
		),
	)
	if err != nil {
		panic(err)
	}
	return conn
}

func NewRbacServiceAdminClient(gc *grpcClient.ClientConn) rbacV1.AdminClient {
	return rbacV1.NewAdminClient(gc)
}

func NewRbacServiceRoleClient(gc *grpcClient.ClientConn) rbacV1.RoleClient {
	return rbacV1.NewRoleClient(gc)
}
