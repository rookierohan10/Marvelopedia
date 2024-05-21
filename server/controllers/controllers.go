package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"server/models"

	"go.uber.org/zap"
)

func HomeController(w http.ResponseWriter, r *http.Request, logger *zap.Logger) {
	logger.Info("Entered into Home Controller Function")

	characters := []string{"iron%20man","captain%20america","thor","hulk","black%20widow","thor","hawkeye"}
	
	type CharacterResponse struct {
		Name string
		Img string
		About string
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
			character.Name= item.Name
			character.About= item.Description
			character.Img= item.Thumbnail.Path
		}

		characterSet = append(characterSet, character)
	}

	w.Header().Add("Content-Type","application/json")
	json.NewEncoder(w).Encode(characterSet)
}
