package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

func main() {

	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	//os.Stdout.Write(b)  //print
	fmt.Println(PrintBeautify(group)) //beatuify print

}
func PrintBeautify(v interface{}) string {
	b, err := jsoniter.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("%v", v)
	}
	return string(b)
}

func Json(v interface{}) (string, error) {
	b, err := jsoniter.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
func JsonBeautify(v interface{}) (string, error) {
	b, err := jsoniter.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
