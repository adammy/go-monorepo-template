package httpapi

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig[Config comparable](path, app string) (*Config, error) {
	var config Config

	viper.AddConfigPath(path)
	viper.SetConfigName(app)
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config failed: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("viper unmarshal failed: %w", err)
	}

	return &config, nil
}
