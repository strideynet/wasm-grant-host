package modulesdk

import (
	"io"
	"log/slog"
	"os"

	"google.golang.org/protobuf/proto"

	"github.com/strideynet/wasm-grant-host/modulesdk/types"
)

type HandleFunc func(*types.Request, *slog.Logger) (*types.Response, error)

func Handle(f HandleFunc) {
	reqBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	req := &types.Request{}
	if err := proto.Unmarshal(reqBytes, req); err != nil {
		panic(err)
	}
	res, err := f(req, slog.Default())
	if err != nil {
		panic(err)
	}
	resBytes, err := proto.Marshal(res)
	if err != nil {
		panic(err)
	}
	if _, err := os.Stdout.Write(resBytes); err != nil {
		panic(err)
	}
}
