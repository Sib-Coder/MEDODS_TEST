package service

import "awesomeProject13/internal/model"

type Database interface {
	UpdateRefresh(idx string, reftoken string) (bool, error)
	SelectInfoUser(idx string) (model.Traning, error)
}
type Service struct {
	db Database
}

func New(db Database) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) UpdateRefToken(idx string, reft string) (bool, error) {
	res, err := s.db.UpdateRefresh(idx, reft)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (s *Service) SelectSecret(idx string) (model.Traning, error) {
	res, err := s.db.SelectInfoUser(idx)
	if err != nil {
		return model.Traning{}, err
	}
	return res, nil
}
