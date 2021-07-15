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

var MYNID string = "6c66980770a523a111ccb5e713398c492afcfa10"

var RE_JOIN_DHT_INTERVAL int = 10

func Entropy(len int) []byte {
	var buf bytes.Buffer
	for i := 0; i < len; i++ {
		buf.WriteByte(byte(rand.Intn(256)))
	}
	return buf.Bytes()
}

func InngerGetNeighbor(target []byte, nid []byte, end int) []byte {
	return append(target[:end], nid[end:]...)
}

func GetNeighbor(target []byte, nid []byte) []byte {
	return InngerGetNeighbor(target, nid, 10)
}

type Node struct {
	NID  []byte
	Host string
	Port int
}

var BOOTSTRAP_NODES []Node = []Node{
	{Host: "router.bittorrent.com", Port: 6881},
	{Host: "dht.transmissionbt.com", Port: 6881},
}

func RandomID() []byte {
	s := sha1.New()
	s.Write(Entropy(20))
	return s.Sum(nil)
}

type DHTServer struct {
	IP      net.IP
	Port    int
	NID     []byte
	UDPConn *net.UDPConn
	Nodes   chan Node
}

func NewDHTServer(ip net.IP, port int) *DHTServer {
	nid, _ := hex.DecodeString(MYNID)

	s := &DHTServer{
		IP: ip, Port: port, NID: nid,
	}
	s.Nodes = make(chan Node)
	return s
}

func (s *DHTServer) ServeForever() {
	var err error
	s.UDPConn, err = net.ListenUDP("udp", &net.UDPAddr{
		IP:   s.IP,
		Port: s.Port,
	})

	if err != nil {
		log.Fatal("Listen failed", err)
		panic(err)
	}

	for {
		var data [2048]byte
		n, addr, err := s.UDPConn.ReadFromUDP(data[:])

		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Printf("Addr:%s, data:%v count:%d \n", addr, string(data[:n]), n)
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

func (s *DHTServer) SendFindNode(node *Node) {
	tid := Entropy(2)

	req := FindNodeReq{
		TransactionID: string(tid),
		Type:          "q",
		FuncName:      "find_node",
		Argument: FindNodeReqArgument{
			ID:     string(s.NID),
			Target: string(RandomID()),
		},
	}

	fmt.Println(req, node)

	// s.Nodes <- *node
	// s.SendKRPC(node, req)
}

func (s *DHTServer) SendKRPC(node *Node, v interface{}) {
	content, _ := Marshal(v)

	fmt.Printf("send to : %s data: %T\n", node.Host, content)

}

func (s *DHTServer) JoinDHT() {
	for index, node := range BOOTSTRAP_NODES {
		fmt.Printf("join dht ------%d-----------------\n", index)
		s.SendFindNode(&node)
		fmt.Printf("join dht ------%d------end--------\n", index)
	}
}

func (s *DHTServer) ReJoinDHT() {
	// 加入DHT网络
	for {
		if len(s.Nodes) == 0 {
			s.JoinDHT()
		}
		// 10s
		time.Sleep(time.Second * time.Duration(RE_JOIN_DHT_INTERVAL))
	}
}

func (s *DHTServer) AutoSendFindNode() {
	// 自动发现

	for {
		select {
		case node := <-s.Nodes:
			fmt.Println("auto send find nodes -----------------")
			s.SendFindNode(&node)
			fmt.Println("auto send find nodes ---------end-----")
		}
		time.Sleep(time.Second)
	}

}

func main() {
	server := NewDHTServer(net.IPv4(0, 0, 0, 0), 9090)

	go server.AutoSendFindNode()

	go server.ReJoinDHT()

	// 监听 0.0.0.0:9090
	go server.ServeForever()

	select {}
}
