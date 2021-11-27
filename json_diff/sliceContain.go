package main

import (
	"fmt"
)

type Universe map[string]bool

func NewUniverse(s []string) Universe {
	u := make(Universe)
	for _, i := range s {
		u[i] = true
	}
	return u
}

func (u Universe) CountainSet(s []string) bool {
	for _, i := range s {
		if !u[i] {
			return false
		}
	}
	return true
}

func findSliceDiff(srcSet, subSet []string) (notFoundSet []string) {
	type Universe map[string]bool

	u := make(Universe)
	for _, i := range srcSet {
		u[i] = true
	}
	for _, i := range subSet {
		if !u[i] {
			notFoundSet = append(notFoundSet, i)
		}
	}
	return notFoundSet
}

func main() {
	fmt.Println(NewUniverse([]string{"a", "b", "c"}).CountainSet([]string{"a", "c"}))
	notFoundSet := findSliceDiff([]string{"a", "b", "c"}, []string{"a", "d"})
	if len(notFoundSet) > 0 {
		fmt.Printf("found not in srcSet :%v", notFoundSet)
	} else {
		fmt.Println("found all in srcSet")
	}
}
