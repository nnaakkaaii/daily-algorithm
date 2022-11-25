package pkg

type List interface {
	Push(x interface{})
	Pop() interface{}
	IsEmpty() bool
}
