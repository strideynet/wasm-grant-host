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

Problems with WASM:

- WASI is not stabilized, there will continue to be significant changes in this
  space.
- Current lack of support for concurrency within instantiated WASM module.

Thinking on what would happen if this went further:

- Pull wasm from OCI
- Explore re-using the instantiated WASM module for multiple requests.
  - Pros:
    - Less time spent "instantiating" the module
    - Ability to have "shared state" between executions
  - Cons:
    - "Shared state" between executions 
- Is using main() for the command entrypoint the right strategy?
  - We could allow main() to be for setup, and then a separate exported func
    is called for each request. This also provides an option for versioning
    the API or allowing for multiple RPCs.
- Multiple "RPCs"
  - An arg to main could name the RPC that we wish to invoke or an arg within
    the request could name the RPC.
  - OR: Literally open some kind of bidirectional socket and use gRPC 
  - OR: Stop using main() and invoke exported funcs for each RPC
  - How does this fit in with re-use of instantiated module.
- Concurrency ???
  - If we want to re-use an instantiated module, we'll need to have a pool since
    access may not be thread safe.
  - If we don't re-use instantiated modules, this is a non-problem, we just
    instantiate one for each request.