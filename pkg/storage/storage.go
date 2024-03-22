package storage

import (
	"time"
)

func (s *Storage) Set(key string, value interface{}, ttl time.Duration) {
	if ttl < 0 {
		return
	}
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	expiresAt := time.Now().Add(ttl)
	s.Store[key] = Data{
		Value:     value,
		ExpiresAt: &expiresAt,
	}
	if ttl > 0 {
		time.AfterFunc(ttl, func() {
			s.Delete(key)
		})
	}
}

func (s *Storage) Get(key string) (interface{}, bool) {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	item, ok := s.Store[key]
	if !ok {
		return nil, false
	}
	return item.Value, true
}

func (s *Storage) Delete(key string) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	delete(s.Store, key)
}
