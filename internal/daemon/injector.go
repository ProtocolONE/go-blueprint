// +build wireinject

package daemon

import (
	"context"

	"github.com/google/wire"
	"github.com/ProtocolONE/go-core/v2/pkg/config"
	"github.com/ProtocolONE/go-core/v2/pkg/invoker"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
)

// Build
func Build(ctx context.Context, initial config.Initial, observer invoker.Observer) (*Daemon, func(), error) {
	panic(wire.Build(WireSet, provider.Set, wire.Struct(new(provider.AwareSet), "*")))
}

// BuildTest
func BuildTest(ctx context.Context, initial config.Initial, observer invoker.Observer) (*Daemon, func(), error) {
	panic(wire.Build(WireTestSet, provider.Set, wire.Struct(new(provider.AwareSet), "*")))
}
