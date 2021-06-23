package models

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type City struct {
	ID         string `json:"id"`
	PROVINCEID string `json:"province_id"`
	Name       string `json:"name"`
}
