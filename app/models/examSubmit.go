package models

// ExamSubmit
//
//	@Description: 提交的数据模型
type ExamSubmit struct {
	// 试卷ID
	PaperID string
	// 作答人姓名
	Name string
	// 作答人学号
	UID string
	// 成绩
	Score string
	// 作答次数
	Num int
}
