package submitController

import (
	"QA-System-Server/app/apiExpection"
	"QA-System-Server/app/models"
	"QA-System-Server/app/services/submitService"
	"fmt"
	"log"
	"math"
	"strconv"
)

// SubmitData
//
//	@Description: 提交答题结果的Controller
//	@param ID 试卷ID
//	@param name 答题人姓名
//	@param UID 答题人学号
//	@param score 本次答题成绩
//	@return error 可能的异常
func SubmitData(ID, name, UID, score string) error {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	submit, e := submitService.FetchSubmit(ID, UID)
	if e != nil {
		log.Println("fetch submit table error")
		return e
	}
	if submit.UID == UID {
		if submit.Num > 1 {
			return apiExpection.ReSubmit
		}
		scoreOld, _ := strconv.ParseFloat(submit.Score, 64)
		scoreNew, _ := strconv.ParseFloat(score, 64)
		maxScore := fmt.Sprintf("%.2f", math.Max(scoreOld, scoreNew))
		err := submitService.UpdateSubmit(models.Submit{
			PaperID: ID,
			Name:    name,
			UID:     UID,
			Score:   maxScore,
			Num:     2})
		if err != nil {
			log.Println("create table submit error")
			return err
		}
		return nil
	}

	err := submitService.CreateSubmit(models.Submit{
		PaperID: ID,
		Name:    name,
		UID:     UID,
		Score:   score,
		Num:     1})
	if err != nil {
		log.Println("create table submit error")
		return err
	}
	return nil
}
