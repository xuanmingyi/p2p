package main

import (
	"fmt"
	"log"
	"net"
)

/*
type PingReq struct {
	ID  int64  `bcode:"id"`
	Req string `bcode:"req"`
}
*/

func Server() {
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})

	if err != nil {
		log.Fatal("Listen failed", err)
		panic(err)
	}
	for {
		var data [1024]byte
		n, addr, err := udpConn.ReadFromUDP(data[:])

		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Printf("Addr:%s,data:%v count:%d \n", addr, string(data[:n]), n)
	}

	fmt.Println("server")
}

func Start() {
	fmt.Println("hello world!!")
}

func main() {
	go Server()

	go Start()

	select {}
}
