package window

import (
	"sync"
	"example.com/go43/quota"
)

type Store struct {
	mu sync.RWMutex
	items map[string]quota.QuotaGrant
}

func New() *Store { return &Store{items: make(map[string]quota.QuotaGrant)} }

func (s *Store) Save(e quota.QuotaGrant) {
	s.mu.Lock(); defer s.mu.Unlock()
	s.items[e.ID] = e.Clone()
}

func (s *Store) Get(id string) (quota.QuotaGrant, bool) {
	s.mu.RLock(); defer s.mu.RUnlock()
	e, ok := s.items[id]
	return e.Clone(), ok
}
