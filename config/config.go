package config

type Config struct {
	Image         string
	DBName        string
	User          string
	Password      string
	ContainerPort string
	DB            DB
}
type DB int

const (
	Mysql DB = iota
	Mongo
	Redis
)

type Builder interface {
	BuildEnv() ([]string, error)
	BuildURI(hostIP, hostPort string) (string, error)
}
