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
	err := s.db.Create(&session).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	result := s.db.Where("token = ?", token).Delete(&model.Session{})
	if result.Error != nil {
		return result.Error
	}

	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		return errors.New("No session found with the specified token")
	}

	return nil
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	result := s.db.Model(&model.Session{}).Where("username = ?", session.Username).Updates(map[string]interface{}{
		"token":  session.Token,
		"expiry": session.Expiry,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *sessionsRepoImpl) SessionAvailName(name string) error {
	var session model.Session
	err := s.db.Where("username = ?", name).First(&session).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("No session found with the specified username")
		}
		return err
	}
	return nil
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session
	err := s.db.Where("token = ?", token).First(&session).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.Session{}, errors.New("No session found with the specified token")
		}
		return model.Session{}, err
	}
	return session, nil
}
