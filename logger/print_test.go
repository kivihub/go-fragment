package logger

import (
	"fmt"
	"testing"
)

func Test_PrintStruct(t *testing.T) {
	person := Person{
		name: "kivi",
		age:  18,
	}
	//fmt.Printf("%s", person)
	fmt.Printf("%v", person)
	fmt.Printf("%+v", person)
}

type Person struct {
	name string
	age  int32
}
