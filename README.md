# windrose-go

This runs an HTTP server that responds to requests with an SVG Windrose.

## Build & Run

Build & run the program in one command:
```sh
go run cmd/serve.go
```

Use `ko` to build a container image and push to a registry:

```sh
GITHUB_TOKEN=topsecret
KO_DOCKER_REPO=ghcr.io/user/repo
ko build cmd/serve.go
```

Run the `ko` built image, without pushing to a registry:

```sh
podman run -p8080:8080 -p40000:40000 $(ko build --local cmd/serve)
```

## Usage

The server listens on port 8080.

It has a single endpoint: `/windrose`, which accepts one of two GET parameters: `angle`, or `direction`.

`angle` would be the direction/angle in degrees (0-360)

`direction` is a cardinal direction: N, S, E, W

For example:

```
curl http://localhost:8080/windrose?angle=270
```

```
curl http://localhost:8080/windrose?direction=W
```

## Debugging

If you build the program with `ko`, you can debug in VSCode as follows:

Configure your launch configs in VSCode:

```json
      {
        "name": "Go dlv-dap Remote",
        "type": "go",
        "debugAdapter": "dlv-dap",
        "request": "attach",
        "mode": "remote",
        "port": 40000,
        "host": "127.0.0.1",
        "substitutePath": [
          {
            "from": "${workspaceFolder}",
            "to": "github.com/jknutson/${workspaceFolderBasename}"
          },
          {
            "from": "${env:GOROOT}",
            "to": "."
          }
        ]
      },
```

Run `ko build`, and run the resulting image via podman.
Expose the app port (8080) and the debugging port (40000)
```bash
$ podman run -p8080:8080 -p40000:40000 $(ko build --local --debug cmd/serve.go)
```

You should see output similar to below.

```sh
API server listening at: [::]:40000
2025-02-23T20:44:47Z warn layer=rpc Listening for remote connections (connections are not authenticated nor encrypted)
2025-02-23T20:44:47Z info layer=debugger launching process with args: [/ko-app/serve.go]
2025-02-23T20:44:47Z debug layer=debugger Adding target 9 "/ko-app/serve.go"
```

You can then attach in VSCode, or via dlv CLI.

### dlv CLI

Connect to remote debugger

```sh
$ dlv connect :40000
```

The debugger will attach to the paused program.

Set a breakpoint with the `break` command, then tell the program to `continue` running.

Debugging output and interactions will occur in this terminal.

```sh
(dlv) break github.com/jknutson/windrose-go/windrose.GenWindrose
Breakpoint 1 set at 0x8ffc56 for github.com/jknutson/windrose-go/windrose.GenWindrose() github.com/jknutson/windrose-go/windrose/windrose.go:24
(dlv) continue

```

Note that the program "hangs" and does not return you to a new prompt. The program is now runnning, when you hit e.g. a breakpoint, you will see output here and a `(dlv)` prompt will appear for you to interact with.

Make a request to the application, you should see some output in the dlv terminal, and then be returned to `(dlv)` prompt:

```sh
> [Breakpoint 1] github.com/jknutson/windrose-go/windrose.GenWindrose() github.com/jknutson/windrose-go/windrose/windrose.go:24 (hits goroutine(22):1 total:1) (PC: 0x8ffc56)
(dlv)
```

At this point, you can start your debugging.

If you wanted to check what arguments (args) were passed to the function:

```sh
(dlv) args
angleDeg = 0
svgWindroseBuf = (*bytes.Buffer)(0xc0000bf770)
```

Check out the [Delve Getting Started Guide](https://github.com/go-delve/delve/blob/master/Documentation/cli/getting_started.md) for more Delve information.