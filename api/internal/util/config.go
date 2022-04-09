package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	MysqlHost     string `mapstructure:"MYSQL_HOST"`
	MysqlPort     string `mapstructure:"MYSQL_PORT"`
	MysqlUser     string `mapstructure:"MYSQL_USER"`
	MysqlPassword string `mapstructure:"MYSQL_PASSWORD"`
	MysqlDbName   string `mapstructure:"MYSQL_DBNAME"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DBSource      string
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	config.DBSource = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.MysqlUser, config.MysqlPassword,
		config.MysqlHost, config.MysqlPort, config.MysqlDbName)

	return
}
