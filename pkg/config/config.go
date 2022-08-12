package config

import "github.com/spf13/viper"

type Config struct {
	Port      string `mapstructure:"PORT"`
	EmailHost string `mapstructure:"EMAIL_HOST"`
	EmailPort string `mapstructure:"EMAIL_PORT"`
	Sender    string `mapstructure:"EMAIL_SENDER"`
	Password  string `mapstructure:"G_APP_PASSWORD"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile(".env")

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
