package controllers

import (
	"net/http"

	"backend-go/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "API ONLINE")

}