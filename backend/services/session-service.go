package services

import "carrmod/backend/domain/models"

type SessionManager struct {
	Repo *models.SessionRepo
}

func NewSessionManager(repo *models.SessionRepo) *SessionManager {
	return &SessionManager{repo}
}

func (s SessionManager) NewSession(session models.Session) error {
	return s.Repo.SaveNewSession(session)
}

func (s SessionManager) RemoveSession(email string) int64 {
	return s.Repo.DeleteSessions(email)
}
