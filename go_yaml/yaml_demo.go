package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"gopkg.in/yaml.v2"
)

const MaxDepth = 10

type TestCase struct {
	Name      string      `yaml:"name"`
	Condition interface{} `yaml:"condition,omitempty"`
	Then      interface{} `yaml:"then"`
}

type YAMLData struct {
	ToRawRequestField       string     `yaml:"to_raw_request_field"`
	FromShippingRequestFiel []string   `yaml:"from_shipping_request_fields"`
	TestCase                []TestCase `yaml:"test_case"`
}

func flattenJSON(prefix string, data interface{}, depth int) (map[string]interface{}, error) {
	if depth > MaxDepth {
		return nil, fmt.Errorf("深度大于[%d]，异常", MaxDepth)
	}

	result := make(map[string]interface{})

	switch value := data.(type) {
	case map[string]interface{}:
		for k, v := range value {
			key := fmt.Sprintf("%s.%s", prefix, k)
			res, err := flattenJSON(key, v, depth+1)
			if err != nil {
				return nil, err
			}
			for k, v := range res {
				result[k] = v
			}
		}
	case []interface{}:
		for i, v := range value {
			key := fmt.Sprintf("%s[%d]", prefix, i)
			res, err := flattenJSON(key, v, depth+1)
			if err != nil {
				return nil, err
			}
			for k, v := range res {
				result[k] = v
			}
		}
	default:
		result[prefix] = value
	}

	return result, nil
}

func formatValue(value interface{}) interface{} {
	rv := reflect.ValueOf(value)

	switch rv.Kind() {
	case reflect.String:
		return rv.String()
	case reflect.Bool:
		return rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint()
	case reflect.Float32, reflect.Float64:
		return rv.Float()
	default:
		return value
	}
}

func generateYAML(data map[string]interface{}) string {
	yamlData := YAMLData{
		TestCase: make([]TestCase, 0),
	}

	for key, value := range data {
		if key == "to_raw_request_field" {
			yamlData.ToRawRequestField = value.(string)
		} else {
			yamlData.TestCase = append(yamlData.TestCase, TestCase{
				Name:      key,
				Condition: formatValue(value),
				Then:      formatValue(value),
			})
		}
	}

	yamlBytes, err := yaml.Marshal(&yamlData)
	if err != nil {
		log.Println("Error generating YAML:", err)
		return ""
	}

	return string(yamlBytes)
}

func main() {
	jsonData := `{
		"stringField": "Hello World",
		"intField": 42,
		"floatField": 3.14,
		"boolField": true,
		"nestedField": {
			"nestedString": "Nested Value",
			"nestedInt": 123,
			"nestedBool": false,
			"nestedArray": [1, 2, 3],
			"nestedObject": {
				"objField1": "ObjValue1",
				"objField2": "ObjValue2"
			}
		},
		"arrayField": [1, "two", 3.0, true]
	}`

	var data interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	result, err := flattenJSON("", data, 1)
	if err != nil {
		fmt.Println("Error flattening JSON:", err)
		return
	}

	yamlString := generateYAML(result)
	if yamlString != "" {
		err = ioutil.WriteFile("output.yaml", []byte(yamlString), 0644)
		if err != nil {
			log.Println("Error writing YAML file:", err)
		} else {
			fmt.Println("YAML file generated successfully.")
		}
	}
}
