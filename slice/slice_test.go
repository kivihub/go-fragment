package slice

import (
	"fmt"
	"testing"
)

func TestSliceLen(t *testing.T) {
	var slice []string
	fmt.Println(len(slice))
	slice = append(slice, "1")
	slice = append(slice, "1")
	slice = append(slice, "1")
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func TestSlice(t *testing.T) {
	slice := []string{"1", "3"}
	fmt.Println("size:", len(slice))
	for i, item := range slice {
		fmt.Println("i", i, "item", item)
	}
	fmt.Println("end")
}
