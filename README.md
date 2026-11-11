# CLIProxyAPI Extra Agents plugin

Adds exta agents to CLIProxyAPI as a native plugin, so the main source stays untouched.

- `model_registrar`: registers `claude-mythos-5`, `claude-mythos-5-1` and `claude-mythos-5-2` under the `claude` provider. They show up for Claude credentials.
- `request_interceptor`: for Claude requests to Mythos 5.1 and 5.2 adds the `per-turn-control` beta; for 5.2 also the `mid-conversation-system-clear-at` beta, the `thinking-binding-controls` beta (adaptive thinking) and `thinking.display: "updates"` when the caller did not set one. The betas go in the body `betas` array, which the Claude executor lifts into `Anthropic-Beta`.

## Build

Needs CGO. Expects the CLIProxyAPI checkout at `../CLIProxyAPI` (see `replace` in `go.mod`).

    go test ./...
    go build -buildmode=c-shared -o bin/extraagents.dll .   # .so on Linux, .dylib on macOS

## Install

Copy the library to `plugins/<GOOS>/<GOARCH>/extraagents.<ext>` and enable it:

    plugins:
      enabled: true
      configs:
        extraagents:
          enabled: true
