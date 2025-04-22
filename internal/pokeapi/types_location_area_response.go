package pokeapi

type LocationAreaListResponse struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type LocationAreaResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	GameIndex int    `json:"game_index"`

	EncounterMethodRates []struct {
		EncounterMethod NamedAPIResource `json:"encounter_method"`
		VersionDetails  []struct {
			Rate    int              `json:"rate"`
			Version NamedAPIResource `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`

	Location NamedAPIResource `json:"location"`

	Names []struct {
		Language NamedAPIResource `json:"language"`
		Name     string           `json:"name"`
	} `json:"names"`

	PokemonEncounters []struct {
		Pokemon        NamedAPIResource `json:"pokemon"`
		VersionDetails []struct {
			Version          NamedAPIResource `json:"version"`
			MaxChance        int              `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int                `json:"min_level"`
				MaxLevel        int                `json:"max_level"`
				ConditionValues []NamedAPIResource `json:"condition_values"`
				Chance          int                `json:"chance"`
				Method          NamedAPIResource   `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
