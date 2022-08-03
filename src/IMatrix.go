package mango

type matrix interface {
	Shape() []int
	Print()
	Add(value) value
}

func Shape(m matrix) []int {
	return m.Shape()
}
