package env

import (
	"github.com/wind-coco/testing-in-docker/config"
	"github.com/wind-coco/testing-in-docker/env/mongo"
	"github.com/wind-coco/testing-in-docker/env/mysql"
	"github.com/wind-coco/testing-in-docker/env/redis"
)

type Builder struct {
	config.Builder
}

func NewBuilder(cfg *config.Config) *Builder {
	var envBuilder config.Builder

	switch cfg.DB {
	case config.Mysql:
		envBuilder = mysql.NewBuilder(cfg)
	case config.Mongo:
		envBuilder = mongo.NewBuilder(cfg)
	case config.Redis:
		envBuilder = redis.NewBuilder(cfg)
	}
	return &Builder{
		envBuilder,
	}
}

func (b *Builder) BuildEnv() ([]string, error) {
	return b.Builder.BuildEnv()
}
func (b *Builder) BuildURI(host, port string) (string, error) {
	return b.Builder.BuildURI(host, port)
}
