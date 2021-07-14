package main

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Name string `json:"aaa"`
	Age  int    `json:""age`
}

func TestRandomID(te *testing.T) {
	i := User{
		Name: "xxx", Age: 30,
	}

	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	for n := 0; n < v.NumField(); n++ {
		fmt.Println(t.Field(n), v.Field(n))
	}
}
