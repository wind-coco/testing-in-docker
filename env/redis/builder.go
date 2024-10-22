package redis

import (
	"fmt"
	"github.com/wind-coco/testing-in-docker/config"
)

const (
	redisConnStr string = "redis://%s:%s@%s:%s/0?protocol=3"
)

type builder struct {
	config *config.Config
}

func NewBuilder(config *config.Config) *builder {
	return &builder{
		config: config,
	}
}

func (b *builder) BuildEnv() ([]string, error) {
	return nil, nil
}

func (b *builder) BuildURI(hostIP, hostPort string) (string, error) {
	return fmt.Sprintf(redisConnStr, b.config.User, b.config.Password, hostIP, hostPort), nil
}
