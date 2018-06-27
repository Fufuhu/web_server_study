package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type URLTarget struct {
	URL    string `yaml:"URL"`
	Method string `yaml:"Method"`
}

func main() {
	fmt.Println("test")

	buf, err := ioutil.ReadFile("./conf/client.yaml")

	if err != nil {
		panic(err)
	}

	fmt.Printf("buf: \n%+v\n", string(buf))

	var t URLTarget
	err = yaml.Unmarshal(buf, &t)
	if err != nil {
		panic(err)
	}

	fmt.Printf("t: %+v\n", t)
}
