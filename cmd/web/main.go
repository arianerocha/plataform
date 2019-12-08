package main

import "github.com/krakenlab/plataform/internal/servers/web"

func main() {
	server := web.NewServer()
	server.Setup()

	server.Run()
}
