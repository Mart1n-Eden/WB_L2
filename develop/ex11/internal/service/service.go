package service

import (
	"ex11/internal/model"
	"ex11/internal/storage"
	"fmt"
	"time"
)

type Service struct {
	store *storage.Storage
}

func NewService(s *storage.Storage) *Service {
	return &Service{store: s}
}

func (s *Service) Create(id int, event model.Event) error {
	if !event.Date.After(time.Now()) {
		return fmt.Errorf("past date")
	}

	s.store.Create(id, event)

	return nil
}

func (s *Service) Update(id int, event model.Event) error {
	if !event.Date.After(time.Now()) {
		return fmt.Errorf("past date")
	}

	if err := s.store.Update(id, event); err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(id int) error {
	if err := s.store.Delete(id); err != nil {
		return err
	}

	return nil
}

func (s *Service) EventsForDay(date time.Time) ([]int, error) {
	res, err := s.store.EventsForDay(date)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) EventsForWeek(date time.Time) ([]int, error) {
	res, err := s.store.EventsForWeek(date)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) EventsForMonth(date time.Time) ([]int, error) {
	res, err := s.store.EventsForMonth(date)

	if err != nil {
		return nil, err
	}

	return res, nil
}
