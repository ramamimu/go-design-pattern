package main

import "fmt"

type Phone struct {
	camera    string
	ram       int
	storage   int
	processor string
}

type PhoneFunc func(*Phone)

func (p *Phone) defaultConfig() {
	p.camera = "20px"
	p.ram = 1024
	p.storage = 64 * 1024
	p.processor = "helios"
}

func NewPhone(fn ...PhoneFunc) Phone {
	p := Phone{}
	p.defaultConfig()

	for _, i := range fn {
		i(&p)
	}

	return p
}

func NewCamera(camera string) PhoneFunc {
	return func(p *Phone) {
		p.camera = camera
	}
}

func NewStorage(storage int) PhoneFunc {
	return func(p *Phone) {
		p.storage = storage
	}
}

func NewRam(ram int) PhoneFunc {
	return func(p *Phone) {
		p.ram = ram
	}
}

func main() {
	p := NewPhone(NewCamera("10px"), NewStorage(100))
	q := NewPhone(NewRam(8*1024), NewStorage(90))
	fmt.Println(p)
	fmt.Println(q)
}
