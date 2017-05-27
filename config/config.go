package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/redite/kleng/utils"
)

// LoadConfig provide routine for loading configuration
func LoadConfig() (Config, error) {
	var c Config
	err := godotenv.Load()
	if err != nil {
		return c, err
	}

	cwd, err := utils.GetCWD()
	if err != nil {
		fmt.Printf("Error getting current working directory. %s\n", err.Error())
		return c, err
	}

	file, err := ioutil.ReadFile(fmt.Sprintf("%s/config.json", cwd))
	if err != nil {
		fmt.Printf("Error opening %s. %s\n", fmt.Sprintf("%s/config.json", cwd), err.Error())
	}

	err = json.Unmarshal(file, &c)
	if err != nil {
		fmt.Printf("Error marshalling %s. %s", fmt.Sprintf("%s/config.json", cwd), err.Error())
		return c, err
	}

	err = envconfig.Process("kleng", &c)
	return c, err
}
