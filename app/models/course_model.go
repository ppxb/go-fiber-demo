package models

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Type           int       `json:"type" gorm:"comment:课程种类 1-私教 2-小班课"`
	StartTime      TimeStamp `json:"startTime" gorm:"comment:开始时间"`
	EndTime        TimeStamp `json:"endTime" gorm:"comment:结束时间"`
	MaxPerson      int       `json:"maxPerson" gorm:"comment:课程最大人数"`
	MinPerson      int       `json:"minPerson" gorm:"default:1;comment:课程最小人数"`
	OrderStartTime TimeStamp `json:"orderStartTime" gorm:"comment:预约开始时间"`
}
