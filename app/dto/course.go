package dto

import "go-fiber-demo/app/models"

type CourseAddDto struct {
	ID             uint             `json:"id"`
	Type           int              `json:"type"`
	StartTime      models.TimeStamp `json:"startTime"`
	EndTime        models.TimeStamp `json:"endTime"`
	MaxPerson      int              `json:"maxPerson"`
	MinPerson      int              `json:"minPerson"`
	OrderStartTime models.TimeStamp `json:"orderStartTime"`
}
