package queries

import (
	"errors"
	"go-fiber-demo/app/models"
	"go-fiber-demo/pkg/database"
)

func GetCourses() ([]models.Course, error) {
	var courses []models.Course

	err := database.DB.Table("tb_courses").Find(&courses).Error
	if err != nil {
		return nil, errors.New("databse internal error")
	}

	return courses, nil
}

func AddCourse(course *models.Course) error {
	err := database.DB.Create(&course).Error
	if err != nil {
		return err
	}
	return nil
}
