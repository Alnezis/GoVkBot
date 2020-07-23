package object

type Coordinates struct {
	Latitude  int `json:"latitude"`
	Longitude int `json:"longitude"`
}

type Place struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Latitude  int    `json:"latitude"`
	Longitude int    `json:"longitude"`
	CreatedAt int    `json:"created"`
	Icon      string `json:"icon"`
	Country   string `json:"country"`
	City      string `json:"city"`
}

type Geo struct {
	Type        string       `json:"type"`
	Coordinates *Coordinates `json:"coordinates"`
	Place       *Place       `json:"place"`
}
