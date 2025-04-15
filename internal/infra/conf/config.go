package conf

import "github.com/spf13/viper"

type Config struct {
	PgUser string `mapstructure:"POSTGRES_USER"`
	PgPass string `mapstructure:"POSTGRES_PASSWORD"`
	PgHost string `mapstructure:"POSTGRES_HOST"`
	PgPort string `mapstructure:"POSTGRES_PORT"`
	PgDB   string `mapstructure:"POSTGRES_DB"`
}

func Load(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
