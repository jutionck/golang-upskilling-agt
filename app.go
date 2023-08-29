package main

import "github.com/jutionck/golang-upskilling-agt/delivery"

// var log = logrus.New()

// otomatis terbaca
// func init() {
// 	log.SetFormatter(&logrus.JSONFormatter{})
// 	log.SetLevel(logrus.DebugLevel)
// }

func main() {
	delivery.NewServer().Run()

	// buat file
	// file, err := os.OpenFile("logger.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Info("Failed to log to file!")
	// } else {
	// 	log.Out = file
	// 	defer file.Close()
	// }

	//
	// log.WithFields(logrus.Fields{
	// 	"project": "Golang Upskilling",
	// }).Info("Error message")

	// log.WithFields(logrus.Fields{
	// 	"project": "Golang Upskilling",
	// }).Warn("Warning message")

	// // fake error
	// err = errors.New("this is error message")
	// if err != nil {
	// 	log.WithFields(logrus.Fields{
	// 		"project": "Golang Upskilling",
	// 	}).Error("Error message")
	// }
}

// Native (package diluar built-in)
// 1. Library Driver Database => Postgres => lib/pq atau bisa menggunakan pgx
// 2. Library Godotenv atau Viper => Environment Configuration
// 3. ID => md5 atau hash 2893927-8u38e3nd0-dndd3 => google/uuid
// 4. Password Hashing => bcrypt

// SQLC
// 1. Library Driver Database => Postgres => lib/pq atau bisa menggunakan pgx
// 2. Library Godotenv atau Viper => Environment Configuration
// 3. ID => md5 atau hash 2893927-8u38e3nd0-dndd3 => google/uuid
// 4. Password Hashing => bcrypt

// GORM
// 1. Driver Database => GORM
// 2. Library Godotenv atau Viper => Environment Configuration
// 3. ID => md5 atau hash 2893927-8u38e3nd0-dndd3 => google/uuid
// 4. Password Hashing => bcrypt
