package models

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Artists struct {
	Id            int                 `json:"id"`
	Image         string              `json:"image"`
	Name          string              `json:"name"`
	Members       []string            `json:"members"`
	CreationDate  int32               `json:"creationDate"`
	FirstAlbum    string              `json:"firstAlbum"`
	Locations     string              `json:"locations"`
	ConcertDates  string              `json:"concertDates"`
	Relations     string              `json:"relations"`
	DatesLocation map[string][]string `json:"datesLocations"`
}
