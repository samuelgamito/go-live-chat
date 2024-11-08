package main

import (
	"go-live-chat/cmd/modules"
	"go-live-chat/internal/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			modules.HandlersModule,
			fx.Invoke(
				bootstrap.RegisterHooks,
			),
		),
	).Run()
}
