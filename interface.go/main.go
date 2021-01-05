package main

import "fmt"

type Intrumenter interface {
	Play()
}

type Guitar struct {
	Strings int
}

type Piano struct {
	Keys int
}

func (g Guitar) Play() {
	fmt.Printf("Tzzzzzouinnnnnnng with %d strings \n", g.Strings)
}

func (p Piano) Play() {
	fmt.Printf("tutotutitu with %d keys \n", p.Keys)
}

func main() {
	var instr Intrumenter
	instr = &Guitar{5}
	instr.Play()

	instr = &Piano{45}
	instr.Play()
}
