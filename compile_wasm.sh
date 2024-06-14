cd wasm_lib
GOOS=js GOARCH=wasm go build -o ../static/wasm/main.wasm .
