package main

import (
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared/utils"
	"gopkg.in/yaml.v2"
)

func loadConfig(file string) (*shared.YamlConfigure, error) {
	var cfg shared.YamlConfigure
	fileCtn, err := utils.ReadFileFromSource(file)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(fileCtn, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
