package dispatcher

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ProtocolONE/go-blueprint/internal/dispatcher/common"
	"github.com/ProtocolONE/go-blueprint/pkg/graphql"
	"github.com/ProtocolONE/go-core/v2/pkg/invoker"
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
)

// Dispatcher
type Dispatcher struct {
	ctx    context.Context
	cfg    Config
	appSet AppSet
	provider.LMT
}

// dispatch
func (d *Dispatcher) Dispatch(echoHttp *echo.Echo) error {
	// middlewares:
	for _, h := range d.appSet.GraphQL.Middleware() {
		echoHttp.Use(echo.WrapMiddleware(h))
	}
	// middleware#2: recover
	echoHttp.Use(middleware.Recover())
	// middleware#1: CORS
	if d.cfg.Debug {
		echoHttp.Use(middleware.CORS())
	} else {
		echoHttp.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			//AllowOrigins:     d.cfg.Cors.Allowed,
			AllowMethods:     []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
			AllowHeaders:     []string{"*"},
			AllowCredentials: false,
		}))
	}

	d.appSet.GraphQL.Routers(echoHttp)

	return nil
}

// Config
type Config struct {
	Debug   bool `fallback:"shared.debug"`
	WorkDir string
	invoker *invoker.Invoker
}

// OnReload
func (c *Config) OnReload(callback func(ctx context.Context)) {
	c.invoker.OnReload(callback)
}

// Reload
func (c *Config) Reload(ctx context.Context) {
	c.invoker.Reload(ctx)
}

type AppSet struct {
	GraphQL *graphql.GraphQL
}

// New
func New(ctx context.Context, set provider.AwareSet, appSet AppSet, cfg *Config) *Dispatcher {
	set.Logger = set.Logger.WithFields(logger.Fields{"service": common.Prefix})
	return &Dispatcher{
		ctx:    ctx,
		cfg:    *cfg,
		appSet: appSet,
		LMT:    &set,
	}
}
