package main

import (
	"fmt"
	"net"
	"testing"
)

func TestRandomID(te *testing.T) {
	conn, err := net.Dial("udp", "router.bittorrent.com:6881")

	conn.Write([]byte("test1111"))

	fmt.Println(conn, err)
}
