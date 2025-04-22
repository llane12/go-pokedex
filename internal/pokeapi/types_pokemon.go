package pokeapi

type Pokemon struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"Height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		Stat     NamedAPIResource `json:"stat"`
		Effort   int              `json:"effort"`
		BaseStat int              `json:"base_stat"`
	} `json:"stats"`
	Types []struct {
		Slot int              `json:"slot"`
		Type NamedAPIResource `json:"type"`
	} `json:"types"`
}
