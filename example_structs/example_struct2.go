package examplestructs

type rectange struct {
	width, height float32
}

type circle struct {
	radis float32
}

func (r rectange) area() float32 {
	return r.height * r.width
}

func (r rectange) perimeter() float32 {
	return r.height*2 + r.width*2
}

func (r circle) area() float32 {
	return 2 * 3.14 * r.radis * r.radis
}

func (r circle) perimeter() float32 {
	return 2 * 3.14 * r.radis
}

func InitExampleStruct2() {
	rectange := rectange{
		height: 2,
		width:  4,
	}
	measure(rectange, "rectange")

	circle := circle{
		radis: 10,
	}
	measure(circle, "circle")

}
