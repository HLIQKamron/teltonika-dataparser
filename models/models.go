package models

type Device struct {
	IMEI string `json:"imei"`
	Lon  string `json:"lon"`
	Lat  string `json:"lat"`
}
