package main

import (
	"context"
	"fmt"
	"log"

	jd "github.com/josephburnett/jd/lib"
)

// getJSONNode 如果s不是json，自动补充成json
func getJSONNode(ctx context.Context, s string) (jd.JsonNode, error) {
	jsonNode, err := jd.ReadJsonString(s)
	if err != nil {
		log.Printf("ReadJsonString err: %v, field value: %s", err, s)
		jsonNode, err = jd.ReadJsonString(fmt.Sprintf(`{"666": "%s"}`, s))
		if err != nil {
			return nil, fmt.Errorf("ReadJsonString err: %v, field value: %s", err, s)
		}
	}

	return jsonNode, nil
}

func getJSONDiff(ctx context.Context, a, b string) (jd.Diff, error) {
	jsonNodeA, err := getJSONNode(ctx, a)
	if err != nil {
		return nil, err
	}

	jsonNodeB, err := getJSONNode(ctx, b)
	if err != nil {
		return nil, err
	}

	return jsonNodeA.Diff(jsonNodeB), nil
}

func main() {
	ctx := context.Background()

	jsonA := `{"name": "John", "age": 30}`
	jsonB := `{"name": "Jane", "age": 25}`

	diff, err := getJSONDiff(ctx, jsonA, jsonB)
	if err != nil {
		log.Fatalf("Error getting JSON diff: %v", err)
	}
	dffStr := diff.Render()
	fmt.Println("The diff between JSON A and JSON B is: \n", dffStr)
}
