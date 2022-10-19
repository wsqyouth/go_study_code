package main

type Callable interface {
	Call() interface{}
	//Back(r interface{})
}
