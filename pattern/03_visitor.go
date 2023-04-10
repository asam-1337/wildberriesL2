package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type shape interface {
	getType() string
	accept(v visitor)
}

type square struct {
	side int
}

func (s *square) getType() string {
	return "square"
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

type triangle struct {
	height int
	side   int
}

func (t *triangle) getType() string {
	return "triangle"
}

func (t *triangle) accept(v visitor) {
	v.visitForTriangle(t)
}

type visitor interface {
	visitForSquare(s *square)
	visitForTriangle(t *triangle)
}

type areaCalculator struct {
	area int
}

func (ac *areaCalculator) visitForSquare(s *square) {
	ac.area = s.side * s.side
	fmt.Println("area of square:", ac.area)
}

func (ac *areaCalculator) visitForTriangle(t *triangle) {
	ac.area = t.height * t.side / 2
	fmt.Println("area of triangle:", ac.area)
}

func main() {
	square := &square{side: 2}
	ac := &areaCalculator{}
	square.accept(ac)
}
