package mango

import (
	"fmt"
	"reflect"
	"testing"
)

func TestShape_01(t *testing.T) {
	shape := []int{10, 10}
	obj := new(denseMatrix)
	obj.shape = shape
	ret := obj.Shape()

	if !reflect.DeepEqual(shape, ret) {
		shapeS := fmt.Sprint(shape)
		retS := fmt.Sprint(ret)
		t.Fatalf(`Shape() = %s, want %s`, shapeS, retS)
	}
}

func TestShape_02(t *testing.T) {
	shape := []int{10, 10}
	obj := new(denseMatrix)
	obj.shape = shape
	ret := Shape(obj)

	if !reflect.DeepEqual(shape, ret) {
		shapeS := fmt.Sprint(shape)
		retS := fmt.Sprint(ret)
		t.Fatalf(`Shape() = %s, want %s`, shapeS, retS)
	}
}

func TestPrint_01(t *testing.T) {

}
