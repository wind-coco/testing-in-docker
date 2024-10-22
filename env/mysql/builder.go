package mysql

import (
	"fmt"
	"github.com/wind-coco/testing-in-docker/config"
)

const mysqlConnStr string = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

type builder struct {
	config *config.Config
}

func NewBuilder(config *config.Config) *builder {
	return &builder{
		config: config,
	}
}
func (b *builder) BuildEnv() ([]string, error) {
	var env []string

	envMap := map[string]string{
		"MYSQL_ROOT_PASSWORD": b.config.Password,
		"MYSQL_DATABASE":      b.config.DBName,
	}
	for envKey, envVar := range envMap {
		env = append(env, envKey+"="+envVar)
	}

	return env, nil
}

func (b *builder) BuildURI(hostIP, hostPort string) (string, error) {
	return fmt.Sprintf(mysqlConnStr, b.config.User, b.config.Password, hostIP, hostPort, b.config.DBName), nil
}
