package utils

import (
	"fmt"
	"github.com/lancelrq/go-iris-starter-kit/src/structs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var GlobalConfig structs.SystemConfiguration

func InitGlobal() {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(fmt.Errorf("failed to load configuration: %s", err.Error()))
	}
	err = yaml.Unmarshal(yamlFile, &GlobalConfig)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal configuration: %s", err.Error()))
	}

}