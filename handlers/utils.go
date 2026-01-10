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
func GetArtistDate(dates models.DateIndex , artistId int) (models.Date, bool){
	var Date models.Date
	for _, v := range dates.Index {
		if v.ID == artistId{
			Date = v
			return  Date ,true
		}
	}
	return Date, false
}