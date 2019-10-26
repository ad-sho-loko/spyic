package spyic

import (
	"fmt"
	"reflect"
)

type SliceStringer struct {
	Slice interface{}
}

func (s SliceStringer) String() string {
	slice := reflect.ValueOf(s.Slice)
	return fmt.Sprint(slice)
}
