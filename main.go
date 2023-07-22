package main

import "go-mongo-fiber/cmd/server"

func main() {
	app := server.Server{}
	app.StartServer()
}
