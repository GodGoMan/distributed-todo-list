package main

import "github.com/GodGoMan/dts/dependencies/client"

func main() {
	cln := client.NewClient("192.168.10.39:6969", "vanilla")
	cln.ConnecToServe()
}
