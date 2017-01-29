package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/acoshift/acourse/app"
	"gopkg.in/yaml.v2"
)

func main() {
	config := app.Config{}
	{
		configFile := os.Getenv("CONFIG")
		if len(configFile) == 0 {
			configFile = "config.yaml"
		}
		bs, err := ioutil.ReadFile(configFile)
		if err != nil {
			log.Fatal(err)
		}
		err = yaml.Unmarshal(bs, &config)
		if err != nil {
			log.Fatal(err)
		}
	}
	app.Run(config)
}
