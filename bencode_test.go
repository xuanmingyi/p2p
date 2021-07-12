package main

import (
	"fmt"
	"testing"
)

type T4 struct {
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
	if string(content) != target {
		fmt.Printf("content: %s\ntarget : %s\n", string(content), target)
	}
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

	t4 := T4{
		TransactionID: "aa",
		Type:          "r",
		FuncName:      "get_peers",
		DataString: []string{
			"sss", "bbb", "ccc",
		},
		Args: []Argument{
			{ID: "bbbbb", InfoHash: "bbbbbbbbbbbbbbbbbb"}, {ID: "abcdefg", InfoHash: "aaaaaaaaaaaaaaaaaa"},
		},
	}

	if !Helper(t4, "d1:t2:aa1:y1:r1:q9:get_peers3:fffl3:sss3:bbb3:ccce1:ald2:id5:bbbbb9:info_hash18:bbbbbbbbbbbbbbbbbbed2:id7:abcdefg9:info_hash18:aaaaaaaaaaaaaaaaaaeee") {
		t.Errorf("error data\n")
		return
	}

}
