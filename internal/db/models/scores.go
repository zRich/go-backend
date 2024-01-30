package models

import "gorm.io/gorm"

// Scroe 记录学生完成任务的成绩

type Score struct {
	gorm.Model
	//学生的学号
	StudentNo string  `json:"studentNo,omitempty"`
	Student   Student `gorm:"foreignKey:StudentNo;references:StudentNo" json:"student,omitempty"`
	//任务的编号
	TaskNo string `json:"taskNo,omitempty"`
	Task   Task   `gorm:"foreignKey:TaskNo;references:TaskNo" json:"task,omitempty"`
	//学生的成绩，默认为0
	Score int `gorm:"default:0" json:"score,omitempty"`
	//学生的测试报告ID, has one relationship
	TestReportID uint `json:"testReportID,omitempty"`
	//学生的测试报告
	TestReport TestReport `gorm:"foreignKey:TestReportID;references:ID" json:"testReport,omitempty"`
}

// TestReport, 用于记录学生的测试报告, 作为给学生评分的依据
type TestReport struct {
	gorm.Model
	//学生的学号
	StudentNo string `json:"studentNo,omitempty"`
	//报告内容，以 json 格式存储
	Report string `json:"report,omitempty"`
}
