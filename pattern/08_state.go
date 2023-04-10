package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type state interface {
	doAction()
}

type context struct {
	state state
}

type state1 struct {
}

func (s *state1) doAction() {
	fmt.Println("do action by state 1")
}

type state2 struct {
}

func (s *state2) doAction() {
	fmt.Println("do action by state 2")
}

func main() {
	ctx := context{state: &state1{}}
}