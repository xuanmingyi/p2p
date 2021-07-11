package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func Marshal(d interface{}) (content []byte, err error) {
	t := reflect.TypeOf(d)
	v := reflect.ValueOf(d)
	var buf bytes.Buffer
	buf.Write([]byte("d"))
	for i := 0; i < v.NumField(); i++ {
		t_t := t.Field(i)
		t_v := v.Field(i)

		switch t_v.Kind() {
		case reflect.Invalid:
			fmt.Printf("invalid\n")

		case reflect.Int64:
			// 数字
			tag := t_t.Tag.Get("bcode")
			if tag != "" {
				buf.Write([]byte(fmt.Sprintf("%d:%s", len(tag), tag)))
				buf.Write([]byte(fmt.Sprintf("i%de", t_v.Int())))
			}
		case reflect.String:
			// 字符串
			tag := t_t.Tag.Get("bcode")
			if tag != "" {
				buf.Write([]byte(fmt.Sprintf("%d:%s", len(tag), tag)))
				buf.Write([]byte(fmt.Sprintf("%d:%s", t_v.Len(), t_v)))
			}
		case reflect.Slice:
			// 列表
			tag := t_t.Tag.Get("bcode")
			if tag != "" {
				buf.Write([]byte("l"))
				for i := 0; i < t_v.Len(); i++ {
					b, e := Marshal(v)
					if e != nil {
						return nil, nil
					}
					buf.Write(b)
				}
				buf.Write([]byte("e"))
			}
		case reflect.Struct:
			// 结构
			tag := t_t.Tag.Get("bcode")
			if tag != "" {
				fmt.Println(t_t.Tag)
			}
		}
	}
	buf.Write([]byte("e"))
	return buf.Bytes(), nil
}

func Unmarshal(data []byte, v interface{}) (err error) {
	return nil
}
