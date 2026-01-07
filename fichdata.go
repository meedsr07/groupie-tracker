package main

const (
	ArtistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	LocationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	DatesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	RelationsURL = "https://groupietrackers.herokuapp.com/api/relation"
)

func FitchArtistData() ([]Artist, error) {
	data, err := GetAPIData(ArtistsURL)
	if err != nil {
		return nil, err
	}

	artistsLocal, err := DecodeDataArtiste(data)
	if err != nil {
		return nil, err
	}

	return artistsLocal, nil
}

func FitchLocationData() ([]Location, error) {
	data, err := GetAPIData(LocationsURL)
	if err != nil {
		return nil, err
	}

	locationsLocal, err := DecodeDataLocation(data)
	if err != nil {
		return nil, err
	}

	return locationsLocal, nil
}
