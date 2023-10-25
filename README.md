# go-rest-templ

A simple template for a REST API server in Go.

Feel free to create an issue or a pull request if you feel like it should have more features.

## Run

Run the server with `go run main.go` or build it with `go build` and run the binary.

## Configuration

The server can be configured in the following ways:

- Command line flags

  ```shell
    go run main.go --port=8080 --addr=localhost --cat-facts=false --cat-facts-url="https://catfact.ninja/fact"
  ```

- Configuration file
  ```shell
    go run main.go -c example-config.yaml
  ```

Note that when using the config `-c/--config` flag, none of the other flags can be used. When using a YAML config file, none of the fields have default values, so all of them must be specified.

## Usage

```shell
$ go run main.go -h
Usage:
  go-rest [flags]

Flags:
  -a, --addr string            Address to run the server on (default "localhost")
  -f, --cat-facts              Enable cat facts (default true)
  -u, --cat-facts-url string   URL to get cat facts from (default "https://catfact.ninja/fact")
  -c, --config string          Config file path (cannot use with any other flag)
  -h, --help                   help for go-rest
  -p, --port int               Port to run the server on (default 8080)
```
