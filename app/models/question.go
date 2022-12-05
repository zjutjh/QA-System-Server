package models

// Question
//
//	@Description: 试卷题目模型
type Question struct {
	// 试卷 ID
	ExamId int
	// 题干
	Topic string
	// 题目类型，1为单选，2为多选
	TypeNum int
	// 合并后的选项，用';'间隔
	Options string
	// 合并后的答案，用';'间隔
	Answer string
}
