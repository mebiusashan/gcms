package config

import "github.com/BurntSushi/toml"

type Server struct {
	Port string
	Addr string
}
type Config struct {
	Server Server
}

func Read(path string) (*Config, error) {
	var cfx Config
	_, err := toml.DecodeFile(path, &cfx)
	return &cfx, err
}
