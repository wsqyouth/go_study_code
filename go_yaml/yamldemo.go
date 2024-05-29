package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type TestCase struct {
	Name      string      `yaml:"name"`
	Condition interface{} `yaml:"condition,omitempty"`
	Then      interface{} `yaml:"then"`
}

type YAMLData struct {
	TestCases []TestCase `yaml:"test_cases"`
}

func convertToJSONString(data interface{}) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func generateYAML(testCases []TestCase) ([]byte, error) {
	yamlData := YAMLData{
		TestCases: testCases,
	}

	yamlBytes, err := yaml.Marshal(&yamlData)
	if err != nil {
		return nil, err
	}

	return yamlBytes, nil
}

func main() {
	testCase1 := TestCase{
		Name: "Shipper.AddressLine1 error test 字段类型不一致",
		Condition: map[string]interface{}{
			"shipment": map[string]interface{}{
				"ship_from": map[string]interface{}{
					"street1": "1863",
					"street2": "S. 3850 W.",
					"street3": "Salt Lake City",
				},
				"ship_to": map[string]interface{}{
					"street1": "1863to",
					"street2": "S. 3850 W.to",
					"street3": "Salt Lake City to",
				},
			},
		},
		Then: 1863,
	}

	testCase2 := TestCase{
		Name:      "Some other test case",
		Condition: "FIX ME",
		Then: map[string]interface{}{
			"marlsha": "ObjValue2",
		},
	}

	testCases := []TestCase{testCase1, testCase2}

	for i := range testCases {
		if jsonData, err := convertToJSONString(testCases[i].Then); err == nil {
			testCases[i].Then = jsonData
		} else {
			log.Printf("Error converting to JSON string: %v", err)
		}
	}

	yamlBytes, err := generateYAML(testCases)
	if err != nil {
		log.Printf("Error generating YAML: %v", err)
		return
	}

	err = ioutil.WriteFile("output.yaml", yamlBytes, 0644)
	if err != nil {
		log.Printf("Error writing YAML file: %v", err)
		return
	}

	fmt.Println("YAML file generated successfully.")
}
