package modulesdk

import (
	"encoding/json"
	"io"
	"log/slog"
	"os"

	"github.com/strideynet/wasm-grant-host/modulesdk/types"
)

var Log = slog.With("component", "module")

type HandleFunc func(types.Request) (*types.Response, error)

func Handle(f HandleFunc) {
	reqBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	req := types.Request{}
	if err := json.Unmarshal(reqBytes, &req); err != nil {
		panic(err)
	}

	res, err := f(req)
	if err != nil {
		panic(err)
	}
	resBytes, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	if _, err := os.Stdout.Write(resBytes); err != nil {
		panic(err)
	}
}
