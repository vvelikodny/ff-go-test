package services

import "sync"

// SessionService represents interface of sessionService service
// SessionService registers sessions keys for all requests
type SessionService interface {
	Register(key string) bool
}

// inMemSessionService implements session service to save sessions keys in-memory
type inMemSessionService struct {
	mux  sync.Mutex
	keys map[string]bool
}

// NewInMemSessionService creates new inMemSessionService
func NewInMemSessionService() SessionService {
	return &inMemSessionService{
		keys: make(map[string]bool),
	}
}

// Register registers new key
// Return true if ket registered, otherwise false
func (s *inMemSessionService) Register(key string) bool {
	s.mux.Lock()
	defer s.mux.Unlock()

	_, exists := s.keys[key]
	if !exists {
		s.keys[key] = true
	}

	return !exists
}
