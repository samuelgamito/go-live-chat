package bootstrap

import (
	"context"
	"fmt"
	"go-live-chat/internal/handlers"
	"go.uber.org/fx"
	"net/http"
)

func RegisterHooks(lifecycle fx.Lifecycle, h *handlers.Handler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				fmt.Println("Starting application in :8080")
				go func() {
					err := http.ListenAndServe(":8080", h.Runner)
					if err != nil {
						panic(err)
					}
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application")
				return nil
			},
		},
	)
}
