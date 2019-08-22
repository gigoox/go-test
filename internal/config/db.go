package config

import "github.com/kelseyhightower/envconfig"

type Db struct {
	Driver string `json:"-" envconfig:"GO_DB_DRIVER" default:"mysql"`
	Server string `json:"-" envconfig:"GO_DB_SERVER" default:"localhost"`
	User   string `json:"-" envconfig:"GO_DB_USER" default:"gestor"`
	Pass   string `json:"-" envconfig:"GO_DB_PASS" default:"gestor123"`
	Port   string `json:"-" envconfig:"GO_DB_PORT" default:"3306"`
	Db     string `json:"-" envconfig:"GO_DB_DB" default:"gestordb"`
}

func NewDb() (*Db, error) {
	config := &Db{}
	err := envconfig.Process("", config)
	return config, err
}
