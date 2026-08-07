package main

import (
	"github.com/voska/frescatto"
	"github.com/voska/vtexkit/cli"
)

// version is injected at build time via -ldflags.
var version = "dev"

func main() {
	cli.Main(cli.App{
		Store:       frescatto.Store,
		Version:     version,
		Description: "Frescatto fish & seafood CLI for humans and AI agents.",
	})
}
