package main

type CallableFunc func(v ...interface{}) interface{}

type CallableJob struct {
	Handler CallableFunc
	Params  []interface{}
}
