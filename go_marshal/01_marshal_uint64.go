package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	hybridTrafficSiteList := []uint64{102, 21, 106}

	b, err := json.Marshal(hybridTrafficSiteList)
	if err != nil {
		fmt.Println("error:", err)
	}
	hybridTrafficSiteListStr := string(b)
	fmt.Printf("%T\n",hybridTrafficSiteListStr)
	fmt.Printf("%v\n",hybridTrafficSiteListStr)

}