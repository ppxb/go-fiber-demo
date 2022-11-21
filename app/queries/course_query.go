package queries

import (
	"errors"
	"go-fiber-demo/app/dto"
	"go-fiber-demo/app/models"
	"go-fiber-demo/pkg/database"
	"go-fiber-demo/pkg/utils"
)

func GetCourses() ([]dto.CourseAddDto, error) {
	var res []models.Course
	var courses []dto.CourseAddDto

	err := database.DB.Table("tb_courses").Find(&res).Error
	if err != nil {
		return nil, errors.New("database internal error")
	}

	utils.Struct2StructByJson(res, &courses)
	return courses, nil
}

func AddCourse(course *models.Course) error {
	err := database.DB.Create(&course).Error
	if err != nil {
		return err
	}
	return nil
}
