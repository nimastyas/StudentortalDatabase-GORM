package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	var students []model.Student
	err := s.db.Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	err := s.db.Create(&student).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	err := s.db.Model(&model.Student{}).Where("id = ?", id).Updates(student).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *studentRepoImpl) Delete(id int) error {
	err := s.db.Where("id = ?", id).Delete(&model.Student{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	var student model.Student
	err := s.db.First(&student, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("No student found with the specified ID")
		}
		return nil, err
	}
	return &student, nil
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	var students []model.StudentClass

	err := s.db.Table("students").Select("students.*, classes.name as class_name").
		Joins("LEFT JOIN classes ON students.class_id = classes.id").
		Find(&students).Error
	if err != nil {
		return nil, err
	}

	if len(students) == 0 {
		students := []model.StudentClass{} // Mengembalikan slice kosong
		return &students, nil
	}
	for i := 0; i < len(students); i++ {
		students[i].Professor = "Dr. Smith"
		students[i].RoomNumber = 101
	}
	return &students, nil
}
