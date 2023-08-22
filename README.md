# Golang Upskilling

## Config

Pada file `config.go` terdapat 2 (dua) library env yaitu `godotenv` dan `viper`. Keduanya bisa digunakan tetapi silahkan pilih salah satu.

Jika ingin menggunakan `godotenv` maka pada settingan `viper` silahkan `comment` seperti berikut:

```go
func (c *Config) ReadConfig() error {

	// COMMENT: jika menggunakan godotenv
	// vp := utils.NewViperUtil("environment", "dev", "env")
	// err := vp.LoadEnv()
	// if err != nil {
	// 	return err
	// }

	// c.DBConfig = DBConfig{
	// 	Host:     vp.GetEnv("DB_HOST", "localhost"),
	// 	Port:     vp.GetEnv("DB_PORT", "5432sss"),
	// 	Name:     vp.GetEnv("DB_NAME", "postgres"),
	// 	Password: vp.GetEnv("DB_PASSWORD", "P@ssw0rd"),
	// 	User:     vp.GetEnv("DB_USER", "postgres"),
	// 	Driver:   vp.GetEnv("DB_DRIVER", "postgres"),
	// }

	// UNCOMMENT: jika menggunakan godotenv
	err := godotenv.Load("environment/.env")
	if err != nil {
		return fmt.Errorf("fatal error config file: %w", err)
	}

	c.DBConfig = DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	return nil
}
```

`Note`:

- Pada folder `environment` silahkan buat file baru dengan nama `.env` yang di duplikat dari `.env.example` kemudian isi pada bagian berikut:

```env
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_DRIVER=
```

Jika ingin menggunakan `viper` silahkan comment bagian `godotenv` menjadi berikut:

```go
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
```

`Note`:

- Pada folder `environment` silahkan buat file baru dengan nama `dev.env` yang di duplikat dari `.env.example` kemudian isi pada bagian berikut:

```env
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_DRIVER=
```

## Run

Jika sudah jalankan program dengan perintah berikut:

```bash
go run .
```
