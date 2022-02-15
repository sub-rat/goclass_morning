package examplestructs

import "fmt"

type geometry interface {
	area() float32 // method signatures
	perimeter() float32
}

func measure(g geometry, name string) {
	fmt.Printf("area of %s= %v\n", name, g.area())
	fmt.Printf("perimeter of %s= %v\n", name, g.perimeter())
}
