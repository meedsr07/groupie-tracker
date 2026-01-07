package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetAPIData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func DecodeDataArtiste(data []byte) ([]Artist, error) {
	var artists []Artist
	err := json.Unmarshal(data, &artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}
func DecodeDataLocation( data []byte)([]Location,error){
	var locations []Location
	err := json.Unmarshal(data,&locations)
	if err!=nil{
		return nil,err
	}
	return locations,err
}
