package models

// QuestionView
//
//	@Description: 将题目的选项切片后的题目模型
type QuestionView struct {
	// 题干
	Topic string
	// 题目类型，1为单选，2为多选
	TypeNum int
	Options []string
}
