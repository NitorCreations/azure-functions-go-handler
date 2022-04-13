package function

import (
	"encoding/json"
	"os"
)

func NewConfigFromFile(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	config := Config{}
	err = json.NewDecoder(file).Decode(&config)
	return &config, err
}

func NewConfig(data []byte) (*Config, error) {
	config := &Config{}
	err := json.Unmarshal(data, config)
	return config, err
}
