package main

import (
	"fmt"
	"reflect"
)

type Bird interface {
	Fly()
}

func main() {
	birdType := reflect.TypeOf((*Bird)(nil)).Elem()
	duck := &Duck{}
	duckTypePtr := reflect.TypeOf(duck)
	duckType := reflect.TypeOf(Duck{})

	fmt.Println(duckTypePtr.Implements(birdType))
	fmt.Println(duckType.Implements(birdType))
	duck.Fly()

	ptrSize := 4 << (^uintptr(0) >> 63)
	fmt.Println(uintptr(0))
	fmt.Println(uintptr(0) >> 63)
	fmt.Println(^uintptr(0))
	fmt.Println(^uintptr(0) >> 63)
	fmt.Println(ptrSize)

	fmt.Println(60 << 3)
}
