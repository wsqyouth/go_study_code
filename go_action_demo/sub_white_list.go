package main

import (
	"encoding/json"
	"fmt"

	//"io/ioutil"
	"reflect"
)

/*
主要学习：
1. map和切片的遍历及查找
2. json的打印
3. 文件的写入
*/
func main() {
	mpMap := map[string]uint32{}
	shareMap := map[string]uint32{}
	snsMap := map[string]uint32{}

	fmt.Printf("src len mp :%v, share:%v, sns:%v\n", len(mpMap), len(shareMap), len(snsMap))
	hasAllWhiteConfig := getWhiteCofig()
	fmt.Println(len(hasAllWhiteConfig))
	notFullMpMap := delHasFullElement(mpMap, hasAllWhiteConfig)
	fmt.Println(len(notFullMpMap))
	GenerateDataJSON("mp", notFullMpMap)

	notFullShareMap := delHasFullElement(shareMap, hasAllWhiteConfig)
	fmt.Println(len(notFullShareMap))
	GenerateDataJSON("share", notFullShareMap)

	notFullSnsMap := delHasFullElement(snsMap, hasAllWhiteConfig)
	fmt.Println(len(notFullSnsMap))
	GenerateDataJSON("sns", notFullSnsMap)
}

func delHasFullElement(srcMap map[string]uint32, whiteConfig []string) map[string]uint32 {
	dstMap := make(map[string]uint32)
	for k, v := range srcMap {
		if (v != 0) && !InArray(k, whiteConfig) {
			dstMap[k] = v
		}
	}
	return dstMap
}

func getWhiteCofig() []string {
	return []string{}
}

//InArray 判断是否在数组中
func InArray(needle interface{}, haystack interface{}) (exists bool) {
	exists = false

	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(haystack)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) {
				exists = true
				return exists
			}
		}
	}

	return exists
}

func GenerateDataJSON(name string, srcMap map[string]uint32) error {
	confBytes, err := json.Marshal(srcMap)
	if err != nil {
		return err
	}

	//fileName := name + ".json"
	//err = ioutil.WriteFile(fileName, confBytes, 0644)
	//if err != nil {
	//	return err
	//}
	fmt.Println(string(confBytes))
	return nil
}
