package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"server/models"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func HomeController(w http.ResponseWriter, r *http.Request, logger *zap.Logger) {
	logger.Info("Entered into Home Controller Function")

	characters := []string{"iron%20man","captain%20america","thor","hulk","black%20widow","thor","hawkeye"}
	
	type CharacterResponse struct {
		ID 		int
		Name 	string
		Img 	string
		About 	string
	}

	characterSet :=  []CharacterResponse{}

	for _, item := range characters{
		response, apierr := http.Get(os.Getenv("HOST")+"characters?name="+item+"&ts=1&apikey="+os.Getenv("PUBLIC_KEY")+"&hash="+os.Getenv("HASH_KEY"))
		if apierr != nil{
			logger.Sugar().Errorf("Error in fetching api: %v", apierr.Error())
		}

		responseData, datafetchingErr := io.ReadAll(response.Body)
		if datafetchingErr != nil {
			logger.Sugar().Errorf("Error in reading response data: %v",datafetchingErr.Error())
		}

		var characterObject models.CharacterResponse
		json.Unmarshal(responseData, &characterObject)
		
		var character CharacterResponse
		for _,item := range characterObject.Data.Results {
			character.ID= item.ID
			character.Name= item.Name
			character.About= item.Description
			character.Img= item.Thumbnail.Path
		}

		characterSet = append(characterSet, character)
	}

	w.Header().Add("Content-Type","application/json")
	json.NewEncoder(w).Encode(characterSet)
}

func SpecificCharacter(w http.ResponseWriter, r *http.Request, logger *zap.Logger) {
	logger.Info("Entered into Specific Character Logger Function")
	vars := mux.Vars(r)
	charId := vars["character_id"]
	logger.Sugar().Infof("Character ID fetched: %s",charId)

	urlString := os.Getenv("HOST")+"characters/"+charId+"?ts="+os.Getenv("TS")+"&apikey="+os.Getenv("PUBLIC_KEY")+"&hash="+os.Getenv("HASH_KEY")
	logger.Sugar().Infoln("URL String: %s",urlString)
	response, apierr := http.Get(urlString)
	if apierr != nil {
		logger.Sugar().Errorln("Error in fetching URL: %s"+ urlString)
	}
	
	responseData, dataFetchErr := io.ReadAll(response.Body)
	if dataFetchErr != nil {
		logger.Sugar().Infoln("Error in reading response body: %v",dataFetchErr.Error())
	}

	var character models.CharacterResponse
	json.Unmarshal(responseData,&character)

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(character)
}

func SpecificCharacterSeriesCollection(w http.ResponseWriter, r *http.Request, logger *zap.Logger){
	logger.Info("Entered into Character Specific Series Route")
	vars := mux.Vars(r)
	charId := vars["character_id"]
	logger.Sugar().Infoln("Character ID fetched: %s",charId)

	urlString := os.Getenv("HOST")+"characters/"+charId+"/series?ts="+os.Getenv("TS")+"&apikey="+os.Getenv("PUBLIC_KEY")+"&hash="+os.Getenv("HASH_KEY")
	logger.Sugar().Infoln("URL String: %s",urlString)
	response, apierr := http.Get(urlString)
	if apierr != nil {
		logger.Sugar().Errorln("Error in fetching URL: %v",apierr.Error())
	}

	responseData, dataFetchErr := io.ReadAll(response.Body)
	if dataFetchErr != nil {
		logger.Sugar().Infoln("Error in reading response body: %v",dataFetchErr.Error())
	}

	var series models.CharacterSeriesResponse
	json.Unmarshal(responseData, &series)

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(series)
}

func SpecificCharacterComicCollection(w http.ResponseWriter, r *http.Request, logger *zap.Logger){
	logger.Info("Entered into Character Specific Series Route")
	vars := mux.Vars(r)
	charId := vars["character_id"]
	logger.Sugar().Infoln("Character ID fetched: %s",charId)

	urlString := os.Getenv("HOST")+"characters/"+charId+"/comics?ts="+os.Getenv("TS")+"&apikey="+os.Getenv("PUBLIC_KEY")+"&hash="+os.Getenv("HASH_KEY")
	logger.Sugar().Infoln("URL String: %s",urlString)
	response, apierr := http.Get(urlString)
	if apierr != nil {
		logger.Sugar().Infoln("Error in fetching URL: %v",apierr.Error())
	}

	responseData, dataFetchErr := io.ReadAll(response.Body)
	if dataFetchErr != nil {
		logger.Sugar().Infoln("Error in reading response body: %v",dataFetchErr.Error())
	}

	var comics models.CharacterComicsResponse
	json.Unmarshal(responseData, &comics)

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(comics)
}

func SpecificCharacterEventsCollection(w http.ResponseWriter, r *http.Request, logger *zap.Logger){
	logger.Sugar().Infoln("Entered into Character Specific Comics Route")
	vars := mux.Vars(r)
	charId := vars["character_id"]
	logger.Sugar().Infoln("Character ID fetched: %s",charId)

	urlString := os.Getenv("HOST")+"characters/"+charId+"/events?ts="+os.Getenv("TS")+"&apikey="+os.Getenv("PUBLIC_KEY")+"&hash="+os.Getenv("HASH_KEY")
	logger.Sugar().Infoln("URL String: %s",urlString)
	response, apierr := http.Get(urlString)
	if apierr != nil {
		logger.Sugar().Infoln("Error in fetching URL: %v",apierr.Error())
	}

	responseData, dataFetchErr := io.ReadAll(response.Body)
	if dataFetchErr != nil {
		logger.Sugar().Infoln("Error in reading response body: %v", dataFetchErr.Error())
	}

	var events models.CharacterEventsResponse
	json.Unmarshal(responseData, &events)

	w.Header().Set("Content-Type","applications/json")
	json.NewEncoder(w).Encode(events)
}

func SpecificCharacterStoriesCollection(w http.ResponseWriter, r *http.Request, logger *zap.Logger){
	logger.Sugar().Infoln("Entered into Character Specific Series Route")
	vars := mux.Vars(r)
	charId := vars["character_id"]
	logger.Sugar().Infoln("Character ID fetched: %s",chardId)

	urlString := os.Getenv("HOST")+"characters/"+charId+"/stories?ts="+os.Getenv("TS")+"&apikey="+os.Getenv("PUBLIC_KEY")+"&hash="+os.Getenv("HASH_KEY")
	logger.Sugar().Infoln("URL String: %s",urlString)
	response, apierr := http.Get(urlString)
	if apierr != nil {
		logger.Sugar().Infoln("Error in fetching URL: %v", apierr.Error())
	}

	responseData, dataFetchErr := io.ReadAll(response.Body)
	if dataFetchErr != nil {
		logger.Sugar().Infoln("Error in reading response body: %v", dataFetchErr.Error())
	}

	var stories models.CharacterStoriesResponse
	json.Unmarshal(responseData, &stories)

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(stories)
}