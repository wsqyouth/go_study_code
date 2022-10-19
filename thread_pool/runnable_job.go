package main

type RunnableFunc func(v ...interface{})

type RunnableJob struct {
	Handler RunnableFunc
	Params  []interface{}
}
