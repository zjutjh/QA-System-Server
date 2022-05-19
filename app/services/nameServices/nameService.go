package nameServices

import (
	"QA-System-Server/app/apiExpection"
	"QA-System-Server/app/models"
	"QA-System-Server/config/database"
	"log"
)

func GetName(id string) (*string, error) {
	var name []models.NameMap

	result := database.DB.Where(models.NameMap{
		ID: id,
	}).Find(&name)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		log.Fatal(apiExpection.ParamError)
		return nil, apiExpection.ParamError
	}
	name_ := &name[0].Name
	return name_, nil
}
