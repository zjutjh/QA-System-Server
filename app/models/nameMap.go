package models

import "QA-System-Server/app/utils"

// NameMap
//
//	@Description: 试卷名称模型
type NameMap struct {
	// 试卷或投票 ID
	ID int
	// 试卷或投票名称
	Name string
	// 截止作答时间
	DeadLine utils.LocalTime
}
