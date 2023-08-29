package config

import (
	"fmt"

	"github.com/jutionck/golang-upskilling-agt/utils"
)

type ApiConfig struct {
	ApiHost string
	ApiPort string
}

type DBConfig struct {
	Host     string
	Port     string
	Name     string
	Password string
	User     string
	Driver   string
}

type FileConfig struct {
	Env string
}

type Config struct {
	DBConfig
	ApiConfig
	FileConfig
}

func (c *Config) ReadConfig() error {

	// COMMENT: jika menggunakan godotenv
	vp := utils.NewViperUtil("environment", "dev", "env")
	err := vp.LoadEnv()
	if err != nil {
		return err
	}

	c.DBConfig = DBConfig{
		Host:     vp.GetEnv("DB_HOST", "localhost"),
		Port:     vp.GetEnv("DB_PORT", "5432sss"),
		Name:     vp.GetEnv("DB_NAME", "postgres"),
		Password: vp.GetEnv("DB_PASSWORD", "P@ssw0rd"),
		User:     vp.GetEnv("DB_USER", "postgres"),
		Driver:   vp.GetEnv("DB_DRIVER", "postgres"),
	}

	c.ApiConfig = ApiConfig{
		ApiHost: vp.GetEnv("API_HOST", "localhost"),
		ApiPort: vp.GetEnv("API_PORT", "8888"),
	}

	c.FileConfig = FileConfig{
		Env: vp.GetEnv("MIGRATION", "dev"),
	}

	// UNCOMMENT: jika menggunakan godotenv
	// err := godotenv.Load("environment/.env")
	// if err != nil {
	// 	return fmt.Errorf("fatal error config file: %w", err)
	// }

	// c.DBConfig = DBConfig{
	// 	Host:     os.Getenv("DB_HOST"),
	// 	Port:     os.Getenv("DB_PORT"),
	// 	Name:     os.Getenv("DB_NAME"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	User:     os.Getenv("DB_USER"),
	// 	Driver:   os.Getenv("DB_DRIVER"),
	// }

	if c.DBConfig.Host == "" || c.DBConfig.Port == "" ||
		c.DBConfig.Name == "" || c.DBConfig.User == "" || c.DBConfig.Password == "" ||
		c.ApiConfig.ApiHost == "" || c.ApiConfig.ApiPort == "" || c.FileConfig.Env == "" {
		return fmt.Errorf("missing required environment variables")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.ReadConfig(); err != nil {
		return nil, err
	}
	return cfg, nil
}
