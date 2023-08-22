package config

import "github.com/jutionck/golang-upskilling-agt/utils"

type DBConfig struct {
	Host     string
	Port     string
	Name     string
	Password string
	User     string
	Driver   string
}

type Config struct {
	DBConfig
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

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.ReadConfig(); err != nil {
		return nil, err
	}
	return cfg, nil
}
