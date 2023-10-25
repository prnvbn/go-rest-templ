# go-rest-templ

A simple template for a REST API server in Go.

Feel free to create an issue or a pull request if you feel like it should have more features.

## Usage

```shell
$ go run main.go -h
A simple REST server template

Usage:
  go-rest [flags]

Flags:
  -a, --addr string            Address to run the server on (default "localhost")
  -c, --cat-facts              Enable cat facts (default true)
  -u, --cat-facts-url string   URL to get cat facts from (default "https://catfact.ninja/fact")
  -h, --help                   help for go-rest
  -p, --port int               Port to run the server on (default 8080)
```
