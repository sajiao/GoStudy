package main

import "fmt"

type Duck struct {
	Name string
}

func (this *Duck) Fly() {
	fmt.Println("this duck can fly.")
}
