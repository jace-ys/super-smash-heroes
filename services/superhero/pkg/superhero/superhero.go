package superhero

type Superhero struct {
	ID         int    `json:"id,omitempty"`
	RealName   string `json:"realName,omitempty"`
	AlterEgo   string `json:"alterEgo,omitempty"`
	PowerStats *Stats `json:"powerStats,omitempty"`
}

type Stats struct {
	Intelligence int `json:"intelligence,omitempty"`
	Strength     int `json:"strength,omitempty"`
	Speed        int `json:"speed,omitempty"`
	Durability   int `json:"durability,omitempty"`
	Power        int `json:"power,omitempty"`
	Combat       int `json:"combat,omitempty"`
}
