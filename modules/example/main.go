package main

import (
	"github.com/strideynet/wasm-grant-host/modulesdk"
	"github.com/strideynet/wasm-grant-host/modulesdk/types"
)

func main() {
	modulesdk.Handle(func(req types.Request) (*types.Response, error) {
		if len(req.Target) == 0 {
			return &types.Response{Allow: false}, nil
		}
		return &types.Response{Allow: true}, nil
	})
}
