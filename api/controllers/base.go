package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

// Server struct
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize - all routes initilise
func (server *Server) Initialize() {
	server.Router = mux.NewRouter()
	server.InitialRoutes()
}

// Run
func (server *Server) Run(addr string) {
	fmt.Println("Listen to the port ", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
