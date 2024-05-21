package routers

import (
	"encoding/json"
	"net/http"
	"server/controllers"
	"server/models"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func CreateRouter() models.Router {
	router := models.Router{
		Router: mux.NewRouter(),
	}
	return router
}

func CreateAppRoutes(router models.Router, logger *zap.Logger) {
	r := router.Router

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Fetched Home Route")
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := struct {
			Msg    string
			Status int
		}{
			Msg: "Server running",
			Status: http.StatusOK,
		}
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	r.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Fetched Home Route")
		controllers.HomeController(w, r, logger)
	})
}
