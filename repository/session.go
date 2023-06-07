package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"

	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
	if err := s.db.Create(&session).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	var sss model.Session
	result := s.db.Where("token = ?", token).Delete(&sss)
	if result.RowsAffected == 0 {
		return errors.New("no product updated")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	result := s.db.Where("username = ?", session.Username).Updates(&session)
	if result.RowsAffected == 0 {
		return result.Error
	}
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) SessionAvailName(name string) error {
	var sss model.Session
	result := s.db.Where("username = ?", name).First(&sss)
	if result.Error != nil {
		return result.Error
	}
	return nil
	// TODO: replace this
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	var sss model.Session
	result := s.db.Where("token = ?", token).First(&sss)
	if result.Error != nil {
		return model.Session{}, result.Error
	}
	return sss, nil // TODO: replace this
}
