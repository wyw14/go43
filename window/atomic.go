package window

import (
	"errors"
	"example.com/go43/quota"
)

var ErrNotFound = errors.New("quotawindow item not found")

func (s *Store) Update(id string, fn func(*quota.QuotaGrant) error) error {
	s.mu.Lock(); defer s.mu.Unlock()
	current, ok := s.items[id]
	if !ok { return ErrNotFound }
	work := current.Clone()
	if err := fn(&work); err != nil { s.items[id] = work.Clone(); return err }
	s.items[id] = work.Clone()
	return nil
}
