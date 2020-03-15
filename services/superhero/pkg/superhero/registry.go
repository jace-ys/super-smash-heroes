package superhero

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	pb "github.com/jace-ys/super-smash-heroes/services/superhero/api/superhero"
)

type SuperheroRegistry interface {
	Find(fullName, alterEgo string) (*pb.Superhero, error)
}

type SuperheroAPIClient struct {
	baseURL string
}

type SuperheroAPIClientConfig struct {
	Token string
}

type superheroAPIResponse struct {
	Response string               `json:"response"`
	Error    string               `json:"error,omitempty"`
	Results  []superheroAPIResult `json:"results,omitempty"`
}

type superheroAPIResult struct {
	AlterEgo   string `json:"name"`
	Powerstats struct {
		Intelligence int32 `json:"intelligence,string"`
		Strength     int32 `json:"strength,string"`
		Speed        int32 `json:"speed,string"`
		Durability   int32 `json:"durability,string"`
		Power        int32 `json:"power,string"`
		Combat       int32 `json:"combat,string"`
	} `json:"powerstats"`
	Biography struct {
		FullName string `json:"full-name"`
	} `json:"biography"`
	Image struct {
		URL string `json:"url"`
	} `json:"image"`
}

func NewSuperheroAPIClient(token string) *SuperheroAPIClient {
	return &SuperheroAPIClient{
		baseURL: fmt.Sprintf("https://superheroapi.com/api/%s", token),
	}
}

func (c *SuperheroAPIClient) Find(fullName, alterEgo string) (*pb.Superhero, error) {
	resp, err := http.Get(fmt.Sprintf("%s/search/%s", c.baseURL, alterEgo))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var superheroes superheroAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&superheroes)
	if err != nil {
		return nil, err
	}

	if superheroes.Error == "" {
		for _, superhero := range superheroes.Results {
			if strings.EqualFold(superhero.Biography.FullName, fullName) && strings.EqualFold(superhero.AlterEgo, alterEgo) {
				return &pb.Superhero{
					FullName:     superhero.Biography.FullName,
					AlterEgo:     superhero.AlterEgo,
					ImageUrl:     superhero.Image.URL,
					Combat:       superhero.Powerstats.Combat,
					Durability:   superhero.Powerstats.Durability,
					Intelligence: superhero.Powerstats.Intelligence,
					Power:        superhero.Powerstats.Power,
					Speed:        superhero.Powerstats.Speed,
					Strength:     superhero.Powerstats.Strength,
				}, nil
			}
		}
	}

	return nil, ErrSuperheroInvalid
}
