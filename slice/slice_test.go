package slice

import (
	"testing"
)

func TestSliceLen(t *testing.T) {
	var slice []string
	t.Log("slice:", slice)
	t.Log("slice len:", len(slice))
	t.Log("slice cap:", cap(slice))

	slice = append(slice, "1")
	slice = append(slice, "2")
	slice = append(slice, "3")
	t.Log("slice:", slice)
	t.Log("slice len:", len(slice))
	t.Log("slice cap:", cap(slice))
}

func TestSliceRange(t *testing.T) {
	slice := []string{"hello", "go"}
	t.Log("size:", len(slice))

	for i, item := range slice {
		t.Log("index:", i, "item:", item)
	}
}
