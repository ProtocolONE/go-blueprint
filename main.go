package main

import (
	"github.com/ProtocolONE/go-blueprint/cmd/daemon"
	"github.com/ProtocolONE/go-blueprint/cmd/gateway"
	"github.com/ProtocolONE/go-blueprint/cmd/migrate"
	"github.com/ProtocolONE/go-blueprint/cmd/root"
	"github.com/ProtocolONE/go-blueprint/cmd/version"
)

func main() {
	root.Execute(gateway.Cmd, version.Cmd, migrate.Cmd, daemon.Cmd)
}
