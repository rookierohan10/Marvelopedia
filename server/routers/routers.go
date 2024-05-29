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

	r.HandleFunc("/character/{character_id}",func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Fetched Specific Character Route")
		controllers.SpecificCharacter(w,r,logger)
	})

	r.HandleFunc("/character/{character_id}/series",func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Fetched Character Specific Series route")
		controllers.SpecificCharacterSeriesCollection(w,r,logger)
	})

	r.HandleFunc("/character/{character_id}/comics",func(w http.ResponseWriter, r *http.Request) {
		logger.Sugar().Infoln("Fetched Character Specific Comic Route")
		controllers.SpecificCharacterComicCollection(w,r,logger)
	})
	r.HandleFunc("/character/{character_id}/events",func(w http.ResponseWriter, r *http.Request) {
		logger.Sugar().Infoln("Fetched Character Specific Events Route")
		controllers.SpecificCharacterEventsCollection(w,r,logger)
	})
	r.HandleFunc("/character/{character_id}/stories",func(w http.ResponseWriter, r *http.Request) {
		logger.Sugar().Infoln("Fetched Character Specific Stories Route")
		controllers.SpecificCharacterStoriesCollection(w,r,logger)
	})
}
