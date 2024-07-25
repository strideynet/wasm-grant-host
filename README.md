# wasm-grant-host

```cmd
GOOS=wasip1 GOARCH=wasm go build -o ./modules/example/main.wasm ./modules/example/main.go
```
// https://github.com/tetratelabs/wazero/blob/main/cmd/wazero/wazero.go

modulesdk is a low-level SDK for interacting with the grant evaluation API.

policysdk is a high-level SDK that wraps the module SDK to provide a "fluent"
interface for defining policies.

```sh
protoc -I ./modulesdk/types  --go_out=paths=source_relative:./modulesdk/types ./modulesdk/types/types.proto
```