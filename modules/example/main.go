package main

import (
	"log/slog"

	"github.com/strideynet/wasm-grant-host/modulesdk"
	"github.com/strideynet/wasm-grant-host/modulesdk/types"
)

func main() {
	modulesdk.Handle(func(req *types.Request, log *slog.Logger) (*types.Response, error) {
		if len(req.Target.Name) == 0 {
			return &types.Response{Allow: false}, nil
		}
		log.Info("I am logging from the wasm module", "target", req.Target.Name)
		return &types.Response{Allow: true}, nil
	})
}
