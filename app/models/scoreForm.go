package models

type ScoreForm struct {
	ID  string `json:"id"`
	Ans []struct {
		ID  int   `json:"id"`
		Key []int `json:"key"`
	} `json:"ans"`
}
