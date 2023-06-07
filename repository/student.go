package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"

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
	var std []model.Student
	if err := s.db.Find(&std).Error; err != nil {
		return []model.Student{}, nil
	}
	return std, nil
	// TODO: replace this
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	if err := s.db.Create(student).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	result := s.db.Where("id = ?", id).Updates(student)
	if result.RowsAffected == 0 {
		return errors.New("no product updated")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) Delete(id int) error {
	var std model.Student
	result := s.db.Where("id = ?", id).Delete(&std)
	if result.RowsAffected == 0 {
		return errors.New("no product updated")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	var std model.Student
	result := s.db.Where("id = ?", id).First(&std)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}

		return nil, result.Error
	}
	return &std, nil // TODO: replace this
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	std := []model.StudentClass{}
	result := s.db.Raw("SELECT s.name, s.address, c.name as class_name, c.professor, c.room_number as room_number FROM Students s JOIN classes c ON c.id = s.class_id").Scan(&std)
	if result.Error != nil {
		return &[]model.StudentClass{}, result.Error
	}
	return &std, nil // TODO: replace this
}
