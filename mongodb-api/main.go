package main

import "go-dev/mongodb-api/cmd/server"

func main() {
	app := server.Server{}
	app.StartServer()
}
