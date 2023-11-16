package main

import "github.com/GodGoMan/dts-server/depedencies/server"

func main() {
	srv := server.NewServer("192.168.10.39:6969")
	srv.Start()
}
