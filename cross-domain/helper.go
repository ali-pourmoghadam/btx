package crossdomain

import (
	"fmt"

	"gopkg.in/ini.v1"
)

func LoadConfigFile() (*ini.File, error) {

	cfg, err := ini.Load("/root/btx/config.ini")

	if err != nil {
		return nil, fmt.Errorf("failed to load ini file: %s", err)
	}

	return cfg, nil
}
