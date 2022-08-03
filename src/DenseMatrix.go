package mango

type denseMatrix struct {
	shape  []int
	values []any
}

func (d denseMatrix) Shape() []int {
	return d.shape
}

func (d denseMatrix) Add(v value) value {
	return d
}

func (d denseMatrix) Print() {

}
