package main

import "github.com/jutionck/golang-upskilling-agt/delivery"

func main() {
	delivery.NewServer().Run()
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
