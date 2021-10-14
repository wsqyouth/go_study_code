package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/google/go-cmp/cmp"
)

func sample1() {
	type LinkStruct struct {
		description  string
		LinkNameType string
	}

	type PageStruct struct {
		description string
		LinkList    LinkStruct
	}

	c1 := LinkStruct{description: "查看详情", LinkNameType: "VIEW_DETAILS"}
	c2 := LinkStruct{description: "查看", LinkNameType: "VIEW_DETAILS"}

	p1 := PageStruct{description: "通过 TSA 落地页制作工具生成的非电商类网页", LinkList: c1}
	p2 := PageStruct{description: "通过 TSA落地页制作工具生成的非电商类网页", LinkList: c2}

	//opts := make([]cmp.Option, 0)
	//opts := []cmp.Option{
	//	cmp.AllowUnexported(p1),
	//}
	if diff := cmp.Diff(p1, p2, cmp.AllowUnexported(PageStruct{})); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}

	fmt.Println("done")
}
func sample2() {

	type LinkStruct struct {
		Description  string
		LinkNameType string
	}

	type PageStruct struct {
		Description string
		LinkList    LinkStruct
	}

	c1 := LinkStruct{Description: "查看详情", LinkNameType: "VIEW_DETAILS"}
	c2 := LinkStruct{Description: "查看", LinkNameType: "VIEW_DETAILS"}

	p1 := PageStruct{Description: "通过 TSA 落地页制作工具生成的非电商类网页", LinkList: c1}
	p2 := PageStruct{Description: "通过 TSA落地页制作工具生成的非电商类网页", LinkList: c2}

	opts := []cmp.Option{
		cmpopts.IgnoreFields(PageStruct{}, "Description"),
		cmpopts.IgnoreFields(LinkStruct{}, "Description"),
	}
	if diff := cmp.Diff(p1, p2, opts...); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}

	fmt.Println("done")
}
func main() {
	sample1()
	//sample2()
}

var t fakeT

type fakeT struct{}

func (t fakeT) Errorf(format string, args ...interface{}) { fmt.Printf(format+"\n", args...) }
