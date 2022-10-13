package config

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestConfig(t *testing.T) {

	filepath := "./config.yaml"
	b, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	var c = &Config{}
	err = yaml.Unmarshal(b, c)
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
	//config := loadConfig(filepath)
	//t.Log(config)
}
