package crossdomain

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

func LoadConfigFile() (*ini.File, error) {

	cwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	cfg, err := ini.Load(cwd + "/../../config.ini")

	if err != nil {
		return nil, fmt.Errorf("failed to load ini file: %s", err)
	}

	return cfg, nil
}
