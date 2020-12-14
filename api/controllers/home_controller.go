package controllers

import (
	"net/http"

	"github.com/filippoipp/go_backend/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "API RODOU!!")
}