package memory

import (
	"context"
	"errors"
	"sync"
	"time"

	"pos-go/internal/modules/auth/ports"
)

var ErrStateNotFound = errors.New("auth state not found")

type stateItem struct {
	value     ports.AuthState
	expiresAt time.Time
}

type AuthStateStore struct {
	mu    sync.Mutex
	items map[string]stateItem
}

func NewAuthStateStore() *AuthStateStore {
	return &AuthStateStore{
		items: map[string]stateItem{},
	}
}

func (s *AuthStateStore) Put(ctx context.Context, state string, value ports.AuthState, ttl time.Duration) error {
	_ = ctx

	if state == "" {
		return errors.New("state empty")
	}
	if ttl <= 0 {
		return errors.New("ttl invalid")
	}

	s.mu.Lock()
	s.items[state] = stateItem{
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}
	s.mu.Unlock()

	return nil
}

func (s *AuthStateStore) GetDel(ctx context.Context, state string) (ports.AuthState, error) {
	_ = ctx

	s.mu.Lock()
	item, ok := s.items[state]
	if ok {
		delete(s.items, state)
	}
	s.mu.Unlock()

	if !ok {
		return ports.AuthState{}, ErrStateNotFound
	}
	if time.Now().After(item.expiresAt) {
		return ports.AuthState{}, ErrStateNotFound
	}

	return item.value, nil
}

var _ ports.AuthStateStore = (*AuthStateStore)(nil)
