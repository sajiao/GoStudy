package main

import "reflect"

type Calculate interface {
	Calc() int
}

type Addition struct {
	a int
	b int
}

func (this Addition) Calc() int {
	return this.a + this.b
}

type Subtraction struct {
	a int
	b int
}

func (this Subtraction) Calc() int {
	return this.a - this.b
}

type Operation struct {
	Symbol string
	a      int
	b      int
}

func (this Operation) Calc() int {
	t := reflect.f

}
