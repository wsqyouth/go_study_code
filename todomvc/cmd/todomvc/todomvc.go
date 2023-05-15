package main

import (
	"github.com/spf13/pflag"

	"todomvc/pkg/todomvc"
	"todomvc/pkg/todomvc/conf"
)

var (
	cfg = pflag.StringP("config", "c", "", "config file path.")
)

func main() {
	pflag.Parse()
	// init config
	if err := conf.Init(*cfg); err != nil {
		panic(err)
	}

	todomvc.App = todomvc.New(conf.Conf)

	todomvc.App.Run()

}

func addTwoSum(a int64, b int64) int64 {
	return a + b
}
