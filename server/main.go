package main

import (
	"log"
	"net/http"
	"os"
	"server/routers"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger, loggerErr := zap.NewProduction()
	if loggerErr != nil{
		log.Fatal("Error in initializing Logger")
	} else {
		logger.Info("Project started")
	}
	defer logger.Sync()

	enverr := godotenv.Load()
	if enverr != nil {
		logger.Fatal("Error Loading the Environment Variable")
	}

	port := os.Getenv("PORT")
	if port == "" {
		logger.Fatal("Port address not retrieved")
	}

	router := routers.CreateRouter()
	routers.CreateAppRoutes(router, logger)
	logger.Info("Server running at http://localhost"+port)
	http.ListenAndServe(port, router.Router)
}
