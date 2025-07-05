package config

import (
	"os"
	"errors"
	"gopkg.in/yaml.v3"
)

func LoadConfigFile [T any] (filePath string) (*T, error) {
	var cfg *T
    if _, err := os.Stat(filePath); err == nil {
        data, err := os.ReadFile(filePath)
        if err != nil {
            return nil, errors.New("error opening file: " + filePath)
        }
        if err := yaml.Unmarshal(data, &cfg); err != nil {
            return nil, errors.New("error parsing yaml file")
        }
    } else if !errors.Is(err, os.ErrNotExist) {
        return nil, err
    }

	return cfg, nil
}