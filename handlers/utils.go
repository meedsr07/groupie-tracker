package handlers

import (
	"groupie-tracker/models"
)

func GetArtistLocation(locations models.LocationIndex, artistId int) (models.Location, bool) {
	var locs models.Location
	for _, v := range locations.Index {
		if v.ID == artistId {
			locs = v
			return locs, true
		}
	}
	return locs, false
}
