# Gimme Dat CLI

![Gimme Dat](./images/gimme-dat.gif)

Simple CLI used to get movie recommendations and movie times based on your postal code.  Currently only 
Dutch style postal codes are accepted, but this will be updated over time.

## Build

The CLI uses [Go v1.22](https://tip.golang.org/doc/go1.22), and can easily be built with the following 
commands:

```bash
go get .
go build -v ./...
```

### Usage

To get more information on usage run:

```bash
go build -o gimmedat -v cmd/main.go
./gimmedat --help
```