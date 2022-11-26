package main

import (
	"context"
	"fmt"
	"math"

	jsoniter "github.com/json-iterator/go"
	"github.com/thoas/go-funk"
)

func main() {
	v := []int64{1, 2, 3, 4, 5, 6, 7}
	defaultSize := 100
	goRoutineNum := 4
	lowBound := 2
	ret, _ := getChunks(context.Background(), v, len(v), defaultSize, goRoutineNum, lowBound)
	chunks := ret.([][]int64)
	fmt.Printf("chunks: %+v, chunk.len: %v\n", chunks, len(chunks))
}

func getChunks(ctx context.Context, itemList interface{}, listLen, defaultChunkSize, goRoutineSize, lowBound int) (interface{}, error) {
	chunks := funk.Chunk(itemList, defaultChunkSize)

	mapIntf := make([]interface{}, 0)
	b, err := jsoniter.Marshal(chunks)
	if err != nil {
		return chunks, err
	}

	err = jsoniter.Unmarshal(b, &mapIntf)
	if err != nil {
		return chunks, err
	}
	if listLen > lowBound && len(mapIntf) < goRoutineSize {
		newSize := int(math.Ceil(float64(listLen) / float64(goRoutineSize)))
		fmt.Printf("raw: %v, listLen: %v, goRoutineSize: %v, newSize: %v\n", float64(listLen)/float64(goRoutineSize), listLen, goRoutineSize, newSize)
		chunks = funk.Chunk(itemList, newSize)
	}
	return chunks, nil
}
