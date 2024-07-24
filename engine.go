package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"

	"github.com/strideynet/wasm-grant-host/modulesdk/types"
)

type Engine struct {
	wasmBytes []byte
}

func (e *Engine) Evaluate(ctx context.Context, req types.Request) (*types.Response, error) {

	// These bits could be cached and reused across multiple evaluations
	slog.Info("Starting Runtime standup")
	runtimeStart := time.Now()
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)
	wasi_snapshot_preview1.MustInstantiate(ctx, r)
	slog.Info("Runtime standup completed", "duration", time.Since(runtimeStart))

	slog.Info("Starting Compilation")
	compileStart := time.Now()
	compiledModule, err := r.CompileModule(ctx, e.wasmBytes)
	if err != nil {
		return nil, err
	}
	slog.Info("Compilation completed", "duration", time.Since(compileStart))

	slog.Info("Starting request handling")
	start := time.Now()

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// TODO: Explore message passing with gRPC?? or at least proto defined
	// messages...

	input := bytes.NewReader(reqBytes)
	output := bytes.NewBuffer(nil)
	slog.Info("Starting invocation of wasm")
	wasmStart := time.Now()
	modConfig := wazero.NewModuleConfig().
		WithStdout(output).    // Used to capture response
		WithStderr(os.Stderr). // Redirect to some kind of wrapped logger
		WithStdin(input)       // Used to feed in input

	// Instantiate the guest Wasm into the same runtime. It exports the `add`
	// function, implemented in WebAssembly.
	_, err = r.InstantiateModule(ctx, compiledModule, modConfig)
	if err != nil {
		return nil, err
	}
	slog.Info("Finished invocation of  wasm", "duration", time.Since(wasmStart))

	res := types.Response{}
	err = json.Unmarshal(output.Bytes(), &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	slog.Info("Finished request handling", "duration", time.Since(start))
	return &res, nil
}

func (e *Engine) Close() {

}

func NewEngine(byteCode []byte) *Engine {
	return &Engine{
		wasmBytes: byteCode,
	}
}
