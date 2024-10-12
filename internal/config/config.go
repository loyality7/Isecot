package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ScanTimeout int `json:"scan_timeout"`
	PortsToScan []int `json:"ports_to_scan"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
