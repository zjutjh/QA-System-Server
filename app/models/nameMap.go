package models

import "QA-System-Server/app/utils"

// ExamNameMap
//
//	@Description: 试卷名称模型
type ExamNameMap struct {
	// 试卷 ID
	ExamID string
	// 试卷名称
	ExamName string
	// 截止作答时间
	DeadLine utils.LocalTime
}
