package testingindocker

import (
	"os"
	"testing"

	"github.com/wind-coco/testing-in-docker/config"
)

func TestRun(t *testing.T) {

}
func TestMain(m *testing.M) {
	os.Exit(RunInDocker(m, &config.Config{
		Image:         "mysql:5.6",
		User:          "root",
		Password:      "123456",
		DBName:        "test_db",
		DB:            config.Mysql,
		ContainerPort: "3306/tcp",
	}))
}
