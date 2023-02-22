package configs

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type DBConfigs struct {
	User     string `env:"DB_USER,required" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD,required"`
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required" envDefault:"5432"`
	Database string `env:"DB_NAME,required" envDefault:"postgres"`
}

func ParseDBConfigs() (DBConfigs, error) {
	if err := godotenv.Load(); err != nil {
		return DBConfigs{}, err
	}

	var configs DBConfigs
	err := env.Parse(&configs)

	return configs, err
}

func (c *DBConfigs) PGConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.User, c.Password, c.Host, c.Port, c.Database)
}
