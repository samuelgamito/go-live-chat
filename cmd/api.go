package main

import (
	"go-live-chat/cmd/modules"
	"go-live-chat/internal/bootstrap"
	"go-live-chat/internal/configs"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			modules.InfraModule,
			modules.HandlersModule,
			fx.Provide(
				configs.NewEnvConfig,
			),
			fx.Invoke(
				bootstrap.RegisterHooks,
			),
		),
	).Run()
}
