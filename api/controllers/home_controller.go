package controllers

import "net/http"

import "github.com/nayonacademy/goapi/api/responses"

// Home - home responses
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome this awesome API ")
}
