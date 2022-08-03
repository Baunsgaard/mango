package mango

type value interface {
	Print()
}

func Print(v value) {
	v.Print()
}
