package main

type LatLng struct {
	Datetime  int64   `json:"datetime"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
	ALtitude  float64 `json:"alt"`
	HDOP      float64 `json:"hdop"`
	PDOP      float64 `json:"pdop"`
	VDOP      float64 `json:"vdop"`
}
