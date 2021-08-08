package main

import "fmt"

const (
    versionNew uint32 = 1
)

func main() {
    bad_example(versionNew)
    fmt.Printf("\n-------\n")
    good_example(versionNew)
}

func bad_example(version uint32){
    configList := []uint64{1,2,3,4}
    configListNew  := []uint64{5,6,7,8}

    if version == versionNew{
        fmt.Printf("use configList: %v",configListNew)
    }else{
        fmt.Printf("use new  configList: %v",configList)
    }
}

// 换一种写法，使用优先级对其赋值，更优雅
func good_example(version uint32) {
    configList := []uint64{1, 2, 3, 4}

    if version == versionNew {
        configList = []uint64{5, 6, 7, 8}
    }
    fmt.Printf("use configList: %v", configList);
}
