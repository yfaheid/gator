package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL    string `json:"db_url"`
	UserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	var cfg Config
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	path := filepath.Join(homeDir, ".gatorconfig.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (c Config) SetUser(name string) error {
	c.UserName = name
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	path := filepath.Join(homeDir, ".gatorconfig.json")
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
