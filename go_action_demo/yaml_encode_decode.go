package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type AppPlugin struct {
	AppKey    string `yaml:"app_key"`
	AppSecret string `ymal:"app_secret"`
}

func main() {
	writeYaml()
	readYaml()
}

// write users to the YAML file.
func writeYaml() {

	appConfig := map[string]AppPlugin{"app1": {"baidu_app_key", "baidu_app_secret"}}

	data, err := yaml.Marshal(&appConfig)
	if err != nil {
		log.Fatal(err)
	}

	err2 := ioutil.WriteFile("app_config.yaml", data, 0644)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("data written")

}

func readYaml() {
	yfile, err := ioutil.ReadFile("app_config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]AppPlugin)
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		log.Fatal(err2)
	}

	for k, v := range data {
		fmt.Printf("%s: %s\n", k, v)
	}

}
