package submitController

import (
	"QA-System-Server/app/models"
	"QA-System-Server/app/services/submitService"
)

func SubmitData(ID, name, UID, score string) error {
	err := submitService.CreateSubmit(models.Submit{
		ID:    ID,
		Name:  name,
		UID:   UID,
		Score: score})
	return err
}
