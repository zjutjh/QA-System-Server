package models

type ScoreForm struct {
	PaperCode string `json:"paperCode"`
	AnsList   []struct {
		ID  int      `json:"id"`
		Key []string `json:"key"`
	} `json:"ansList"`
}
