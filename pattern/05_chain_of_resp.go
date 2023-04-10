package main

import (
	"log"
	"net/http"
)

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

func AccessMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}
}

type handler interface {
	execute()
	setNextHandler(handler)
}

type LogMiddleware struct {
	next handler
}

func (m *LogMiddleware) execute() {
	log.Println("log middleware is executed")
	if m.next != nil {
		m.next.execute()
	}
}

func (m *LogMiddleware) setNextHandler(h handler) {
	m.next = h
}

type AuthMiddleware struct {
	next handler
}

func (m *AuthMiddleware) execute() {
	log.Println("authorized successfully")
	if m.next != nil {
		m.next.execute()
	}
}

func (m *AuthMiddleware) setNextHandler(h handler) {
	m.next = h
}

type PanicMiddleware struct {
	next handler
}

func (m *PanicMiddleware) execute() {
	log.Println("panic middleware executed")
	defer func() {
		if err := recover(); err != nil {
			log.Println("recovered", err)
		}
	}()

	if m.next != nil {
		m.next.execute()
	}
}

func (m *PanicMiddleware) setNextHandler(h handler) {
	m.next = h
}

func main() {
	p := &PanicMiddleware{}
	a := &AuthMiddleware{}
	l := &LogMiddleware{}
	p.setNextHandler(l)
	l.setNextHandler(a)
	p.execute()
}
