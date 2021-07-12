package main

// https://www.cnblogs.com/LittleHann/p/6180296.html
// https://ops.tips/blog/udp-client-and-server-in-go/#a-udp-server-in-go

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
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

func Entropy(len int) []byte {
	var buf bytes.Buffer
	for i := 0; i < len; i++ {
		buf.WriteByte(byte(rand.Intn(256)))
	}
	return buf.Bytes()
}

func RandomID() string {
	s := sha1.New()
	s.Write(Entropy(20))
	return hex.EncodeToString(s.Sum(nil))
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

func (s *Server) FindNode(node *Node) {
	req := FindNodeReq{
		TransactionID: "aa",
		Type:          "q",
		FuncName:      "find_node",
		Argument: FindNodeReqArgument{
			ID:     "sss",
			Target: "sssss",
		},
	}

	s.SendKRPC(&Nodes[0], req)
}

func (s *Server) SendKRPC(node *Node, v interface{}) {
	fmt.Println(Marshal(v))
}

func (s *Server) Join() {
	for _, node := range Nodes {
		s.FindNode(&node)
	}
}

func (s *Server) ReJoin() {
	// 加入DHT网络
	for {
		s.Join()
		time.Sleep(time.Second * time.Duration(100))
	}
}

func (s *Server) AutoFindNode() {
	// 自动发现

	for {
		time.Sleep(time.Second * time.Duration(100))
	}

}

func main() {
	server := Server{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	}

	go server.Serve()

	go server.ReJoin()

	go server.AutoFindNode()

	select {}
}
