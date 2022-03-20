package services

import (
	"github.com/henilthakor/goapi/domain"
)

func AddNewStudentDetail(s *domain.Student) {
	domain.AddNewStudentDetail(s)
	return
}

func GetStudentDetail(id int) (*domain.Student, error) {
	s, err := domain.GetStudentDetail(id)
	return s, err
}

func UpdateStudentDetail(s *domain.Student) error {
	err := domain.UpdateStudentDetail(s.Id, s)
	return err
}

func DeleteStudentDetail(id int) error {
	err := domain.DeleteStudentDetail(id)
	return err
}
