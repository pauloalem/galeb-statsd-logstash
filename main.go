package main

import (
	"log"
	"net"
	"regexp"
	"strconv"
)

type document struct {
	Client string `json:"client"`
	Metric string `json:"metric"`
	Count  int    `json:"count"`
	App    string `json:"app"`
	Value  int    `json:"value"`
}

func handle(data []byte) (*document, error) {
	r := regexp.MustCompile(`galeb\.(?P<addr>[\w-_]+)\..*.requestTime:(?P<value>\d+)|ms.*`)
	d := r.FindStringSubmatch(string(data))
	value, err := strconv.Atoi(d[2])
	if err != nil {
		return nil, err
	}
	doc := &document{
		Client: "tsuru",
		Metric: "response_time",
		Count:  1,
		App:    d[1],
		Value:  value,
	}
	return doc, nil
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
