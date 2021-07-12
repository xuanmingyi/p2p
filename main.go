package main

// https://www.cnblogs.com/LittleHann/p/6180296.html
// https://ops.tips/blog/udp-client-and-server-in-go/#a-udp-server-in-go

import (
	"fmt"
	"log"
	"net"
)

type Node struct {
	Host string
	Port int
}

var Nodes []Node = []Node{
	{Host: "router.bittorrent.com", Port: 6881},
	{Host: "dht.transmissionbt.com", Port: 6881},
	//{Host: "router.utorrent.com", Port: 6881}, // link error
}

type Server struct {
	IP      net.IP
	Port    int
	UDPConn *net.UDPConn
}

func (s *Server) Serve() {
	var err error
	s.UDPConn, err = net.ListenUDP("udp", &net.UDPAddr{
		IP:   s.IP,
		Port: s.Port,
	})

	if err != nil {
		log.Fatal("Listen failed", err)
		panic(err)
	}

	s.ReJoin()

	for {
		var data [1024]byte
		n, addr, err := s.UDPConn.ReadFromUDP(data[:])

		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Printf("Addr:%s,data:%v count:%d \n", addr, string(data[:n]), n)
	}

}

type FindNodeReqArgument struct {
	ID     string `bcode:"id"`
	Target string `bcode:"target"`
}

type FindNodeReq struct {
	TransactionID string              `bcode:"t"`
	Type          string              `bcode:"y"`
	FuncName      string              `bcode:"q"`
	Argument      FindNodeReqArgument `bcode:"a"`
}

func (s *Server) FindNode() {
	req := FindNodeReq{
		TransactionID: "aa",
		Type:          "q",
		FuncName:      "find_node",
		Argument: FindNodeReqArgument{
			ID:     "sss",
			Target: "sssss",
		},
	}

	s.SendKRPC(req)
}

func (s *Server) SendKRPC(v interface{}) {
	fmt.Println(Marshal(v))
}

func (s *Server) Join() {
	fmt.Println(Nodes)
}

func (s *Server) ReJoin() {
}

func (s *Server) Start() {
	fmt.Println("sniffer start!!")
}

func main() {
	server := Server{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	}

	go server.Serve()

	go server.Start()

	select {}
}
