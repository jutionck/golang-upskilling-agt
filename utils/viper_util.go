package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type ViperUtil interface {
	LoadEnv() error
	GetEnv(key string, defaultValue string) string
}

type viperUtil struct {
	configPath string
	configName string
	configType string
}

func (v *viperUtil) LoadEnv() error {
	viper.AddConfigPath(v.configPath)
	viper.SetConfigName(v.configName)
	viper.SetConfigType(v.configType)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("fatal error config file: %w", err)
	}
	return nil
}

func (v *viperUtil) GetEnv(key string, defaultValue string) string {
	if envVal := viper.GetString(key); len(envVal) != 0 {
		return envVal
	}
	return defaultValue
}

func NewViperUtil(configPath string, configName string, configType string) ViperUtil {
	return &viperUtil{
		configPath: configPath,
		configName: configName,
		configType: configType,
	}
}
