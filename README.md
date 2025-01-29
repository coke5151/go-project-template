# Go (Golang) Project Template

A go (golang) project template with [Air](https://github.com/air-verse/air) (Hot reloading tool) and VSCode settings

This project also turn on some inlay hints in VSCode. You can preview with `cmd/inlayhints/main.go`.

You can also turn it off by changing/deleting the `settings.json` file in `.vscode` folder whenever you want.

## Initial Setup

```bash
go mod init github.com/your-username/your-project-name # initialize the project
go install github.com/air-verse/air@latest # install air
air init # initialize air
```

## Air
Air is a tool for building and reloading your go application.

Your `air.toml` file may look like this:

```toml
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "tmp\\main.exe"
  cmd = "go build -o ./tmp/main.exe ."
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = []
# ...
```

change the `cmd` command to the path you want to run currently.

for example, to run `cmd/inlayhints/main.go`, edit the `.air.toml` file:

```toml
[build]
  cmd = "go build -o ./tmp/main.exe ./cmd/inlayhints/main.go"
```

You only need to run the command `air` in the root of your project to start hot reloading.

