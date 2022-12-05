package examController

import (
	"QA-System-Server/app/apiExpection"
	"QA-System-Server/app/models"
	"QA-System-Server/app/services/nameServices"
	"QA-System-Server/app/services/questionServices"
	"QA-System-Server/app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

// GetExam
//
//	@Description: 获取试卷功能的Controller
//	@param c gin.Context
func GetExam(c *gin.Context) {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	idRaw := c.Query("id")
	timeRaw := c.Query("time")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		log.Println("exam id error")
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}
	time, err := strconv.ParseInt(timeRaw, 10, 64)
	if err != nil {
		log.Println("timestamp error")
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}

	nameMaps, err := nameServices.GetName(id)
	if err != nil {
		log.Println("name_map table error")
		_ = c.AbortWithError(200, apiExpection.ServerError)
		return
	}
	if len(nameMaps) == 0 {
		log.Println("this exam is not existing")
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}
	DeadLine := nameMaps[0].DeadLine.Time
	ExamName := nameMaps[0].Name
	if time > DeadLine.Unix() {
		utils.JsonSuccessResponse(c, gin.H{
			"data": "超出试卷作答时间",
			"name": ExamName,
		}, "EXPIRED")
		return
	}

	questions, err := questionServices.GetQuestions(id)
	if err != nil {
		log.Println("exam table error")
		_ = c.AbortWithError(200, apiExpection.ServerError)
		return
	}

	questionsViews := make([]models.QuestionView, len(questions))
	for i, value := range questions {
		questionsViews[i].Topic = value.Topic
		questionsViews[i].TypeNum = value.TypeNum
		questionsViews[i].Options = strings.Split(value.Options, ";")
	}

	utils.JsonSuccessResponse(c, gin.H{
		"data": questionsViews,
		"name": ExamName,
	}, "SUCCESS")
}

// GetScore
//
//	@Description: 提交答案获取分数的Controller
//	@param c gin.Context
func GetScore(c *gin.Context) {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	var scoreForm models.SubmitForm
	sum := 0

	err := c.ShouldBindJSON(&scoreForm)
	if err != nil {
		log.Println("request form error")
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}

	id, err := strconv.Atoi(scoreForm.ID)
	if err != nil {
		log.Println("exam id error")
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}
	questions, err := questionServices.GetQuestions(id)
	if err != nil {
		log.Println("questions table error")
		_ = c.AbortWithError(200, apiExpection.ServerError)
		return
	}

	if len(questions) != len(scoreForm.Ans) {
		log.Println("answer list length error")
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}
	sort.SliceStable(scoreForm.Ans, func(i, j int) bool {
		return scoreForm.Ans[i].ID < scoreForm.Ans[j].ID
	})
	for i, value := range questions {
		flag := true
		answers := strings.Split(value.Answer, ";")

		if len(answers) != len(scoreForm.Ans[i].Key) {
			continue
		}
		if len(scoreForm.Ans[i].Key) > 1 {
			sort.SliceStable(scoreForm.Ans[i].Key, func(j, k int) bool {
				return scoreForm.Ans[i].Key[j] < scoreForm.Ans[i].Key[k]
			})
		}
		for j := 0; j < len(answers) && flag; j++ {
			answer, _ := strconv.Atoi(answers[j])
			if answer != scoreForm.Ans[i].Key[j] {
				flag = false
				break
			}
		}
		if !flag {
			continue
		}
		sum++
	}
	score := math.Round(((100.0/float64(len(questions)))*float64(sum))*1e2+0.5) * 1e-2

	e := SubmitData(scoreForm.ID,
		scoreForm.Name,
		scoreForm.UID,
		strconv.FormatFloat(score, 'f', 2, 64))
	if e == apiExpection.ReSubmit {
		utils.JsonSuccessResponse(c, "请勿重复提交", "SUCCESS")
	} else if e != nil {
		log.Println("submit table error")
		_ = c.AbortWithError(200, apiExpection.ServerError)
	} else {
		utils.JsonSuccessResponse(c, score, "SUCCESS")
	}
}

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
	submit, e := questionServices.FetchSubmit(ID, UID)
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
		err := questionServices.UpdateSubmit(models.ExamSubmit{
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

	err := questionServices.CreateSubmit(models.ExamSubmit{
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
