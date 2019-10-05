// +build wireinject

package resolver

import (
	"context"

	"github.com/google/wire"
	"github.com/ProtocolONE/go-blueprint/generated/graphql"
	"github.com/ProtocolONE/go-core/v2/pkg/config"
	"github.com/ProtocolONE/go-core/v2/pkg/invoker"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
)

// Build
func Build(ctx context.Context, initial config.Initial, observer invoker.Observer) (graphql.Config, func(), error) {
	panic(wire.Build(provider.Set, WireSet, wire.Struct(new(provider.AwareSet), "*")))
}

// Build
func BuildTest(ctx context.Context, initial config.Initial, observer invoker.Observer) (graphql.Config, func(), error) {
	panic(wire.Build(provider.Set, WireTestSet, wire.Struct(new(provider.AwareSet), "*")))
}
