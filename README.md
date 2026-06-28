# tang-init

A zero-dependency project bootstrapper for Tang Nano FPGAs.

Generates a ready-to-compile Verilog project with a pre-configured Makefile, pin constraints, and blinky code — no memorizing CLI flags or hunting through PDF schematics.

## Usage

Run the binary in an empty directory:

```
./tang-init
```

You'll be prompted for a project name and board selection. The tool creates:

```
<project_name>/
├── Makefile
├── .gitignore
├── constraints/board.cst
├── src/top.v
└── tb/top_tb.v
```

Then `cd <project_name> && make flash` compiles and loads onto your board.

## Supported Boards

| Board         | Chip     | Clock | LEDs |
|---------------|----------|-------|------|
| Tang Nano 9K  | GW1N-9C  | Pin 52| 6    |
| Tang Nano 20K | GW2A-18C | Pin 4 | 6    |

## Build from source

Requires Go 1.21+.

```bash
go build -o tang-init .
```

Cross-compile for other platforms:

```bash
GOOS=windows GOARCH=amd64 go build -o tang-init.exe .
GOOS=darwin  GOARCH=arm64  go build -o tang-init-mac .
GOOS=linux   GOARCH=amd64 go build -o tang-init-linux .
```
