package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	good_example()
	return
}

func bad_example() {
	var grayConf map[string]int
	//旧的配置数据
	grayJsonStr := "{\"GrayID1\":1,\"GrayID2\":1}"
	if err := json.Unmarshal([]byte(grayJsonStr), &grayConf); err != nil {
		panic(err)
	}
	fmt.Printf("gray map is:%+v\n", grayConf)

	//新的配置数据
	newGrayJsonStr := "{\"GrayID1\":1,\"GrayID3\":1}"
	if err := json.Unmarshal([]byte(newGrayJsonStr), &grayConf); err != nil {
		panic(err)
	}
	fmt.Printf("new gray map is:%+v\n", grayConf)
}

func good_example() {
	var grayConf map[string]int
	//旧的配置数据
	grayJsonStr := "{\"GrayID1\":1,\"GrayID2\":1}"
	if err := json.Unmarshal([]byte(grayJsonStr), &grayConf); err != nil {
		panic(err)
	}
	fmt.Printf("gray map is:%+v\n", grayConf)

	//新的配置数据
	newGrayJsonStr := "{\"GrayID1\":1,\"GrayID3\":1}"
	// 方案1： 重置obj对象
	grayConf = nil
	if err := json.Unmarshal([]byte(newGrayJsonStr), &grayConf); err != nil {
		panic(err)
	}
	fmt.Printf("gray map is:%+v\n", grayConf)

	// 方案2：新建一个obj对象
	var newGrayConf map[string]int
	if err := json.Unmarshal([]byte(newGrayJsonStr), &newGrayConf); err != nil {
		panic(err)
	}
	fmt.Printf("new gray map is:%+v\n", newGrayConf)
}

/*
If the map is nil, Unmarshal allocates a new map. Otherwise Unmarshal reuses the existing map, keeping existing entries.
Unmarshal then stores key-value pairs from the JSON object into the map.
The map's key type must either be any string type, an integer, implement json.Unmarshaler, or implement encoding.TextUnmarshaler.
*/
