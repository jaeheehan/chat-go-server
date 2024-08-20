package main

import "chat_server_golang/network"

func main() {
	n := network.NewServer()
	n.StartServer()
}
