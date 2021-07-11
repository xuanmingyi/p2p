package main

import (
	"fmt"
	"testing"
)

type Req struct {
	TransactionID string     `bcode:"t"`
	Type          string     `bcode:"y"`
	FuncName      string     `bcode:"q"`
	DataString    []string   `bcode:"fff"`
	Args          []Argument `bcode:"a"`
}

type Argument struct {
	ID       string `bcode:"id"`
	InfoHash string `bcode:"info_hash"`
}

type T1 struct {
	A string `bcode:"abc"`
}

type T2 struct {
	A int64 `bcode:"def"`
}

type T3 struct {
	A []string `bcode:"ghi"`
}

func Helper(v interface{}, target string) bool {
	content, err := Marshal(v)
	if err != nil {
		return false
	}
	fmt.Println(string(content))
	fmt.Println(target)
	return string(content) == target
}

func TestMarshual(t *testing.T) {
	t1 := T1{
		A: "def",
	}

	if !Helper(t1, "d3:abc3:defe") {
		t.Errorf("error data\n")
		return
	}

	t2 := T2{
		A: 2133,
	}

	if !Helper(t2, "d3:defi2133ee") {
		t.Errorf("error data\n")
		return
	}

	t3 := T3{
		A: []string{
			"aaaaa", "bbbbb", "ccccc",
		},
	}
	if !Helper(t3, "d3:ghil5:aaaaa5:bbbbb5:cccccee") {
		t.Errorf("error data\n")
		return
	}

	/*req := Req{
		TransactionID: "aa",
		Type:          "r",
		FuncName:      "get_peers",
		DataString: []string{
			"sss", "bbb", "ccc",
		},
		Args: []Argument{
			{ID: "abcdefg", InfoHash: "aaaaaaaaaaaaaaaaaa"},
			{ID: "bbbbb", InfoHash: "bbbbbbbbbbbbbbbbbb"},
		},
	}
	content, err := Marshal(req)

	if err != nil {
		t.Errorf("error\n")
	}

	if string(content) != "d1t:ad2:id20:abcdefghij01234567899:info_hash20:mnopqrstuvwxyz123456e1:q9:get_peers1:t2:aa1:y1:qe" {
		fmt.Println(content)
		fmt.Println(string(content))
		t.Errorf("error compare: %s\n", string(content))
	}*/

}
