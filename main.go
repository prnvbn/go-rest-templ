package main

import (
	"go-rest/rest"
)

func main() {
	server := rest.NewServer()
	server.Run()
}
