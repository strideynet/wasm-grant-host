package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/strideynet/wasm-grant-host/modulesdk/types"
)

func main() {
	ctx := context.Background()

	byteCode, err := os.ReadFile("./modules/example/main.wasm")
	if err != nil {
		panic(err)
	}

	eng := NewEngine(byteCode)
	res, err := eng.Evaluate(ctx, types.Request{
		Target: []string{"test"},
	})
	if err != nil {
		panic(err)
	}
	slog.Info("Received response from engine", "response", res)
}
