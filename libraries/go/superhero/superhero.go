package superhero

import (
	"fmt"
	"os"
)

func GetBaseURI() string {
	accessToken, ok := os.LookupEnv("SUPERHERO_API_ACCESS_TOKEN")
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s/%s", "https://superheroapi.com/api", accessToken)
}

type Response struct {
	Response string    `json:"response"`
	Results  []Results `json:"results"`
}

type Results struct {
	ID         string     `json:"id"`
	AlterEgo   string     `json:"name"`
	Powerstats Powerstats `json:"powerstats"`
	Biography  Biography  `json:"biography"`
	Image      Image      `json:"image"`
}

type Powerstats struct {
	Intelligence string `json:"intelligence"`
	Strength     string `json:"strength"`
	Speed        string `json:"speed"`
	Durability   string `json:"durability"`
	Power        string `json:"power"`
	Combat       string `json:"combat"`
}

type Biography struct {
	FullName string `json:"full-name"`
}

type Image struct {
	URL string `json:"url"`
}
