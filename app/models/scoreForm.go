package models

// SubmitForm
//
//	@Description: 提交作答的表单
type SubmitForm struct {
	// 试卷ID
	ID string `json:"id"`
	// 作答人姓名
	Name string `json:"name"`
	// 作答人学号
	UID string `json:"uid"`
	// 选择的结果，题目ID与选择的选项列表成对
	Ans []struct {
		ID  int   `json:"id"`
		Key []int `json:"key"`
	} `json:"ans"`
}
