package main

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Строитель порождающий паттерн, служит для конструирования сложных обьектов

type house struct {
	wallType   string
	floorsType string
	roofType   string
}

type builder interface {
	getHouse() *house
	setWallType()
	setFloorsType()
	setRoofType()
}

type woodenBuilder struct {
	h house
}

func newWoodenBuilder() *brickBuilder {
	return &brickBuilder{
		h: house{},
	}
}

func (b *woodenBuilder) setWallType() {
	b.h.wallType = "wooden"
}

func (b *woodenBuilder) setFloorsType() {
	b.h.floorsType = "linoleum"
}

func (b *woodenBuilder) setRoofType() {
	b.h.wallType = "schist"
}

func (b *woodenBuilder) getHouse() *house {
	return &b.h
}

type brickBuilder struct {
	h house
}

func newBrickBuilder() *brickBuilder {
	return &brickBuilder{
		h: house{},
	}
}

func (b *brickBuilder) setWallType() {
	b.h.wallType = "brick"
}

func (b *brickBuilder) setFloorsType() {
	b.h.floorsType = "parquet"
}

func (b *brickBuilder) setRoofType() {
	b.h.wallType = "metal"
}

func (b *brickBuilder) getHouse() *house {
	return &b.h
}

type director struct {
	b builder
}

func newDirector(builder builder) *director {
	return &director{
		b: builder,
	}
}

func (d *director) buildHouse() *house {
	d.b.setWallType()
	d.b.setFloorsType()
	d.b.setFloorsType()
	return d.b.getHouse()
}
