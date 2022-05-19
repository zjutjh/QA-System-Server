package submitController

import (
	"QA-System-Server/app/apiExpection"
	"QA-System-Server/app/models"
	"QA-System-Server/app/services/submitService"
	"log"
)

func SubmitData(ID, name, UID, score string) error {
	submit, e := submitService.FetchSubmit(UID)
	if e != nil {
		log.Fatal(e)
		return e
	}
	if submit.Name == name {
		return apiExpection.ReSubmit
	}

	err := submitService.CreateSubmit(models.Submit{
		ID:    ID,
		Name:  name,
		UID:   UID,
		Score: score})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
