package main

import (
	"log"
	"week-6-assignment-Furknn/server"
)

func main() {
	svr := server.NewServer()
	err := svr.StartServer(3000)
	if err != nil {
		log.Fatal(err)
	}
}
