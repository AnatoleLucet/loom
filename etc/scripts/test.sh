#!/bin/sh

RUN_STANDARD=false
RUN_WASM=false

while [ $# -gt 0 ]; do
  case "$1" in
    -s|--standard)
      RUN_STANDARD=true
      ;;
    -w|--wasm)
      RUN_WASM=true
      ;;
    --)
      shift
      break
      ;;
    -*)
      echo "Unknown option: $1" >&2
      exit 1
      ;;
    *)
      break
      ;;
  esac
  shift
done

if [ "$RUN_STANDARD" = false ] && [ "$RUN_WASM" = false ]; then
  RUN_STANDARD=true
  RUN_WASM=true
fi

if [ "$RUN_STANDARD" = true ]; then
  echo "Running standard tests..."
  go test -v -race $@
fi

if [ "$RUN_WASM" = true ]; then
  echo "Running wasm tests..."
  env -i \
    HOME="$HOME" \
    PATH="$PATH" \
    GOROOT="$(go env GOROOT)" \
    GOPATH="$(go env GOPATH)" \
    GOCACHE="$(go env GOCACHE)" \
    GOOS=js \
    GOARCH=wasm \
    go test -v -exec="bash $(go env GOROOT)/lib/wasm/go_js_wasm_exec" $@
fi
