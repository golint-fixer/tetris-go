package libs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Env    string `json:"env"`
	Listen struct {
		Address string `json:"address"`
	} `json:"listen"`
	Github struct {
		Token    string `json:"token"`
		Username string `json:"username"`
	} `json:"github"`
	Template struct {
		Path string `json:"path"`
	} `json:"template"`
}

func NewConfig() *Config {
	c := &Config{}
	return c.loadEnvConfig()
}

func NewConfigFile(path string) *Config {
	c := NewConfig()
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)

	err := decoder.Decode(c)
	if err != nil {
		panic(fmt.Sprintf("Error parsing JSON config file: %s", err))
	}

	return c.loadEnvConfig()
}

func (c *Config) loadEnvConfig() *Config {
	if env := os.Getenv("ENV_TETRIS_ENV"); env != "" {
		c.Env = env
	}
	if env := os.Getenv("ENV_TETRIS_GITHUB_TOKEN"); env != "" {
		c.Github.Token = env
	}
	return c
}
