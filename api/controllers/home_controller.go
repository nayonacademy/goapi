package controllers

import (
	"github.com/nayonacademy/goapi/api/responses"
	"net/http"
)

// Home - home responses
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome this awesome API ")
}
