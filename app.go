package main

import (
	"github.com/jutionck/golang-upskilling-agt/delivery"
)

// @title           Upskilling Golang
// @version         1.0
// @description     This is a sample server celler server.

// @host      localhost:8888
// @BasePath  /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @schemes http
func main() {
	delivery.NewServer().Run()
}
