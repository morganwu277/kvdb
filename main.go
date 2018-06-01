package main

import "github.com/morganwu277/kvdb/server"

const (
	port = ":50051"
)

func main() {
	server.StartServer(port)
}
