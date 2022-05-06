package models

type Question struct {
	questionnaireID int
	Stem            string
	TypeNum         int
	Options         string
	answer          string
}
