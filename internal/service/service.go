package service

import (
	"encoding/json"
	"log"
	"main/internal/models"
	"net/http"
	"strings"
)

type Service struct{}

const (
	artistUrl   = "https://groupietrackers.herokuapp.com/api/artists"
	IdArtist    = "https://groupietrackers.herokuapp.com/api/artists/"
	relationUrl = "https://groupietrackers.herokuapp.com/api/relation/"
)

func (s *Service) Allartist() ([]models.Artists, error) {
	var artist []models.Artists

	client := http.Client{}

	res, err := client.Get(artistUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&artist); err != nil {
		return nil, err
	}
	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// jsonErr := json.Unmarshal(body, &artist)
	// if jsonErr != nil {
	// 	log.Fatal(jsonErr)
	// }
	return artist, nil
}

func (s *Service) IdArtist(idurl string) (models.Artists, error) {
	concatUrl := IdArtist + idurl
	var artist models.Artists
	client := http.Client{}
	res1, err := client.Get(concatUrl)
	if err != nil {
		return artist, err
	}
	defer res1.Body.Close()

	if err := json.NewDecoder(res1.Body).Decode(&artist); err != nil {
		return artist, err
	}

	return artist, nil
}

func (s *Service) Relations(idUrl string) (models.Relations, error) {
	concatUrl := relationUrl + idUrl
	var relation models.Relations
	client := http.Client{}
	res, err := client.Get(concatUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&relation); err != nil {
		log.Fatal(err)
	}
	ChangeLocations := ChangeStr(relation.DatesLocations)
	relation.DatesLocations = ChangeLocations
	return relation, nil
}

func ChangeStr(DatesLocation map[string][]string) map[string][]string {
	for i, v := range DatesLocation {
		temp := i
		temp = strings.ReplaceAll(temp, "-", ", ")
		temp = strings.ReplaceAll(temp, "_", " ")
		temp = strings.Title(temp)
		delete(DatesLocation, i)
		DatesLocation[temp] = v
	}
	return DatesLocation
}
