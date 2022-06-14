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

var myName = "kivi"

func TestConvertByValue(t *testing.T) {
	user := User{
		name:     &myName,
		age:      1,
		favorite: "football",
	}
	fmt.Println("user addr", unsafe.Pointer(&user))
	fmt.Println("user_addr", unsafe.Pointer((*User_)(unsafe.Pointer(&user))))
	// go等号复制是浅复制
	user_addr := *(*User_)(unsafe.Pointer(&user))
	fmt.Println("user_addr", unsafe.Pointer(&user_addr))

	// 指针内容更新
	*user.name = "wqw"
	if *user.name == *user_addr.name {
		t.Log("same after update user.name")
	}
	// 非指针内容更新
	user.favorite = "beer"
	if user.favorite != user_addr.favorite {
		t.Log("diff after update user.favorite")
	}
}

func TestConvertByPointer(t *testing.T) {
	user := &User{
		name:     &myName,
		age:      1,
		favorite: "football",
	}
	user_ := (*User_)(unsafe.Pointer(user))

	fmt.Println("user addr", unsafe.Pointer(user))
	fmt.Println("user_ addr", unsafe.Pointer(user_))

	*user.name = "wqw"
	if *user.name == *user_.name {
		t.Log("same after update user.name")
	} else {
		t.Error("diff after update user.name")
	}

	user.favorite = "beer"
	if user.favorite == user_.favorite {
		t.Log("same after update user.favorite")
	} else {
		t.Error("diff after update user.favorite")
	}
}

type User struct {
	name     *string
	age      int8
	favorite string
}

type User_ struct {
	name     *string
	age      int8
	favorite string
}
