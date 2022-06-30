package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	sideLen float32
}

func (s square) area() float32 {
	return s.sideLen * s.sideLen
}

type triangle struct {
	height float32
	base   float32
}

func (t triangle) area() float32 {
	return t.base * t.height / 2
}

type shape interface {
	area() float32
}

type calculator struct {
	total float32
}

func (c calculator) sumAreas(shapes ...shape) float32 {
	var sum float32

	for _, shape := range shapes {
		sum += shape.area()
	}

	return sum
}

//func (c calculator) sumAreas(shapes ...interface{}) float32 {
//	var sum float32
//
//	for _, shape := range shapes {
//		switch shape.(type) {
//		case circle:
//			r := shape.(circle).radius
//			sum += math.Pi * r * r
//		case square:
//			l := shape.(square).sideLen
//			sum += l * l
//		}
//	}
//	return sum
//}

type outPrinter struct {
}

func (op outPrinter) toText(s shape) string {
	return fmt.Sprintf("the area is: %f\n", s.area())
}

func main() {
	c := circle{radius: 5}
	c.area()

	s := square{sideLen: 2}
	s.area()

	t := triangle{height: 10, base: 5}
	s.area()

	out := outPrinter{}
	fmt.Printf(out.toText(c))
	fmt.Printf(out.toText(s))
	fmt.Printf(out.toText(t))

	calc := calculator{}
	//fmt.Println("\ntotal of areas:", calc.sumAreas(c, s))
	fmt.Println("\ntotal of areas:", calc.sumAreas(c, s, t))

}
