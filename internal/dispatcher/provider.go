package dispatcher

import (
	"context"

	"github.com/ProtocolONE/go-blueprint/internal/dispatcher/common"

	"github.com/google/wire"
	"github.com/ProtocolONE/go-core/v2/pkg/config"
	"github.com/ProtocolONE/go-core/v2/pkg/invoker"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
)

// ProviderCfg
func ProviderCfg(cfg config.Configurator) (*Config, func(), error) {
	c := &Config{
		WorkDir: cfg.WorkDir(),
		invoker: invoker.NewInvoker(),
	}
	e := cfg.UnmarshalKeyOnReload(common.UnmarshalKey, c)
	return c, func() {}, e
}

// ProviderDispatcher
func ProviderDispatcher(ctx context.Context, set provider.AwareSet, appSet AppSet, cfg *Config) (*Dispatcher, func(), error) {
	d := New(ctx, set, appSet, cfg)
	return d, func() {}, nil
}

var (
	WireSet = wire.NewSet(
		ProviderDispatcher,
		ProviderCfg,
		wire.Struct(new(AppSet), "*"),
	)

	WireTestSet = wire.NewSet(
		ProviderDispatcher,
		ProviderCfg,
		wire.Struct(new(AppSet), "*"),
	)
)
