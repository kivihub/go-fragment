package unsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestPointer(t *testing.T) {
	s := make([]int, 9, 20)
	//var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + 8))
	fmt.Println(Len, len(s)) // 9 9

	//var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + 16))
	fmt.Println(Cap, cap(s)) // 20 20
}

func TestPointer2(t *testing.T) {
	user := User{
		name:     "kivi",
		age:      1,
		favorite: "football",
	}
	user_ := *(*User_)(unsafe.Pointer(&user))

	fmt.Println("user addr", unsafe.Pointer(&user))
	fmt.Println("user_ addr", unsafe.Pointer(&user_))

	fmt.Println("user", user)
	fmt.Println("user_", user_)

	user.name = "wqw"
	fmt.Println("==update user.name==")
	fmt.Println("user", user)
	fmt.Println("user_", user_)
}

type User struct {
	name     string
	age      int8
	favorite string
}

type User_ struct {
	name     string
	age      int8
	favorite string
}
