package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.String:
		s := v.String()
		fmt.Fprintf(buf, "%d:%s", len(s), s)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "i%de", v.Int())
	case reflect.Slice:
		buf.WriteByte('l')
		for i := 0; i < v.Len(); i++ {
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte('e')
	case reflect.Struct:
		buf.WriteByte('d')
		for i := 0; i < v.NumField(); i++ {
			tag := v.Type().Field(i).Tag.Get("bcode")
			if tag != "" {
				fmt.Fprintf(buf, "%d:%s", len(tag), tag)
				if err := encode(buf, v.Field(i)); err != nil {
					return err
				}
			}
		}
		buf.WriteByte('e')
	}
	return nil
}

func Marshal(d interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(d)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Unmarshal(data []byte, v interface{}) (err error) {
	return nil
}
