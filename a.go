package main

import (
	"encoding/hex"
	"fmt"
	"net"
)

func main() {
	myid := RandomID()
	udpAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}

	socket, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}
	defer socket.Close()

	// write
	dhtaddress, err := net.ResolveUDPAddr("udp", "router.bittorrent.com:6881")
	if err != nil {
		panic(err)
	}
	fmt.Println(dhtaddress)

	tid := Entropy(2)

	req := FindNodeReq{
		TransactionID: string(tid),
		Type:          "q",
		FuncName:      "find_node",
		Argument: FindNodeReqArgument{
			ID:     string(myid),
			Target: string(RandomID()),
		},
	}

	data1, _ := Marshal(req)
	fmt.Println(data1)
	socket.WriteToUDP(data1, dhtaddress)

	// read
	fmt.Println(1111)
	var data []byte = make([]byte, 1024)

	fmt.Println(1112)
	n, addr, err := socket.ReadFromUDP(data)

	fmt.Println(1113)
	if err != nil {
		panic(err)
	}

	fmt.Println(1114)
	fmt.Println(addr, hex.EncodeToString(data[:n]))
}
