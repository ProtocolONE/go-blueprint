package daemon

import (
	"context"

	"github.com/google/wire"
	"github.com/ProtocolONE/go-blueprint/internal/db/domain"
	"github.com/ProtocolONE/go-blueprint/internal/db/repo"
	"github.com/ProtocolONE/go-blueprint/internal/db/trx"
	"github.com/ProtocolONE/go-blueprint/pkg/postgres"
	"github.com/ProtocolONE/go-core/v2/pkg/config"
	"github.com/ProtocolONE/go-core/v2/pkg/invoker"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
)

// Cfg
func Cfg(cfg config.Configurator) (*Config, func(), error) {
	c := &Config{
		invoker: invoker.NewInvoker(),
	}
	e := cfg.UnmarshalKeyOnReload(UnmarshalKey, c)
	return c, func() {}, e
}

// CfgTest
func CfgTest() (*Config, func(), error) {
	c := &Config{
		invoker: invoker.NewInvoker(),
	}
	return c, func() {}, nil
}

// Repo
type Repo struct {
	List domain.ListRepo
}

// Provider
func Provider(ctx context.Context, set provider.AwareSet, appSet AppSet, cfg *Config) (*Daemon, func(), error) {
	g := New(ctx, set, appSet, cfg)
	return g, func() {}, nil
}

var (
	ProviderRepo = wire.NewSet(
		repo.NewListRepo,
		trx.NewTrxManager,
	)
	ProviderRepoProduction = wire.NewSet(
		ProviderRepo,
		wire.Struct(new(Repo), "*"),
		postgres.WireSet,
	)
	ProviderTestRepo = wire.NewSet(
		ProviderRepo,
		wire.Struct(new(Repo), "*"),
		postgres.WireTestSet,
	)
	WireSet = wire.NewSet(
		Provider,
		Cfg,
		ProviderRepoProduction,
		wire.Struct(new(AppSet), "*"),
	)
	WireTestSet = wire.NewSet(
		Provider,
		CfgTest,
		ProviderTestRepo,
		wire.Struct(new(AppSet), "*"),
	)
)
