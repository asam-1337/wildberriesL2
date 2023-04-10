package main

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type algorithm interface {
	sort()
}

type sorter struct {
	sortAlgo algorithm
}

func (s *sorter) setSortAlgo(algo algorithm) {
	s.sortAlgo = algo
}

func (s *sorter) sortSlice() {
	s.sortAlgo.sort()
}

type quickSort struct {
}

func (s *quickSort) sort() {
	fmt.Println("quick sort")
}

type bubbleSort struct {
}

func (s *bubbleSort) sort() {
	fmt.Println("bubble sort")
}

type shellSort struct {
}

func (s *shellSort) sort() {
	fmt.Println("shell sort")
}

func main() {
	s := sorter{}
	s.setSortAlgo(&quickSort{})
	s.sortSlice()

	s.setSortAlgo(&bubbleSort{})
	s.sortSlice()

	s.setSortAlgo(&shellSort{})
	s.sortSlice()
}