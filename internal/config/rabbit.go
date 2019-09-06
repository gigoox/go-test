package config

import "github.com/kelseyhightower/envconfig"

type Rabbit struct {
	Uri string `json:"-" envconfig:"RABBIT_URI" default:"amqp://guest:guest@localhost:5672/"`
	ExchangeName string `json:"-" envconfig:"RABBIT_EXCHANGE_NAME" default:"test-exchange"`
	ExchangeType string `json:"-" envconfig:"RABBIT_EXCHANGE_TYPE" default:"direct"`
	RoutingKey string `json:"-" envconfig:"RABBIT_ROUTING_KEY" default:"test-key"`
	Body string `json:"-" envconfig:"RABBIT_BODY" default:"foobar"`
	Reliable string `json:"-" envconfig:"RABBIT_RELIABLE" default:"false"`
}

func NewRabbit() (*Rabbit, error) {
	rabbit := &Rabbit{}
	err := envconfig.Process("", rabbit)
	return rabbit, err
}

