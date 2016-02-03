package main

import (
	"fmt"
	"log"
	"net"
)

func handle(data []byte) {
	fmt.Println("Received ", string(data))
}

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":8125")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		buf := make([]byte, 1600)
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Print(err)
		}
		handle(buf[0:n])
	}
}
