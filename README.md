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

| Board                | Chip        | Clock   | Freq   | LEDs |
|----------------------|-------------|---------|--------|------|
| Tang Nano (original) | GW1N-1      | Pin 35  | 24 MHz | RGB  |
| Tang Nano 1K         | GW1NZ-1     | Pin 47  | 27 MHz | RGB  |
| Tang Nano 4K         | GW1NSR-4C   | Pin 45  | 27 MHz | 1    |
| Tang Nano 9K         | GW1NR-9C    | Pin 52  | 27 MHz | 6    |
| Tang Nano 20K        | GW2A-18C    | Pin 4   | 27 MHz | 6    |

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
