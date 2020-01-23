package api

import "github.com/nayonacademy/goapi/api/controllers"

var server = controllers.Server{}

// Run service
func Run() {
	server.Initialize()
	server.Run(":8080")
}
