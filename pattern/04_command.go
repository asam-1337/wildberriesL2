package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/
// command
type command interface {
	execute()
}

// command interface implementation
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.turnOn()
}

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.turnOff()
}

type device interface {
	turnOn()
	turnOff()
}

// invoker
type button struct {
	command command
}

func (b *button) on() {
	b.command.execute()
}

func (b *button) off() {
	b.command.execute()
}

type tv struct {
}

func (t *tv) turnOn() {
	fmt.Println("tv is on")
}

func (t *tv) turnOff() {
	fmt.Println("tv is off")
}

func main() {
	tv := &tv{}
	onComm := &onCommand{device: tv}
	offComm := &offCommand{device: tv}
	onButt := button{
		command: onComm,
	}
	offButt := button{
		command: offComm,
	}

	onButt.command.execute()
	offButt.command.execute()
}