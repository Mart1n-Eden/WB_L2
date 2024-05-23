package storage

import (
	"ex11/internal/model"
	"fmt"
	"sync"
	"time"
)

type Storage struct {
	mu     sync.RWMutex
	events map[int]model.Event
}

func NewStorage() *Storage {
	return &Storage{events: make(map[int]model.Event)}
}

func (s *Storage) Create(id int, event model.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.events[id] = event
}

func (s *Storage) Update(id int, event model.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[id]; !ok {
		return fmt.Errorf("no this id")
	}

	s.events[id] = event
	return nil
}

func (s *Storage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[id]; !ok {
		return fmt.Errorf("no this id")
	}

	delete(s.events, id)

	return nil
}

func (s *Storage) EventsForDay(date time.Time) ([]int, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var res []int

	for k, v := range s.events {
		if v.Date == date {
			res = append(res, k)
		}
	}

	return res, nil
}

func (s *Storage) EventsForWeek(date time.Time) ([]int, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var res []int
	targetYear, targetWeek := date.ISOWeek()

	for k, v := range s.events {
		eventYear, eventWeek := v.Date.ISOWeek()
		fmt.Println(eventYear, eventWeek)
		fmt.Println(targetYear, targetWeek)
		if eventYear == targetYear && eventWeek == targetWeek {
			res = append(res, k)
		}
	}

	return res, nil
}

func (s *Storage) EventsForMonth(date time.Time) ([]int, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var res []int

	for k, v := range s.events {
		if v.Date.Month() == date.Month() {
			res = append(res, k)
		}
	}

	return res, nil
}
