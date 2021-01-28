package domain

import (
	"context"
	"errors"
	"sync"
)

type CountryRepository interface {
	FindByName(context.Context, string) (*Country, error)
	SaveOrUpdate(context.Context, *Country) error
}

type InMemoryCountryRepository struct {
	sync.Mutex
	storage map[string]*Country
}

func NewRepository() *InMemoryCountryRepository {
	return &InMemoryCountryRepository{
		storage: make(map[string]*Country),
	}
}

var ErrCountryNotFound = errors.New("specified country does not exist")

func (repo *InMemoryCountryRepository) FindByName(_ context.Context, name string) (*Country, error) {
	repo.Lock()
	defer repo.Unlock()
	c, ok := repo.storage[name]
	if !ok {
		return nil, ErrCountryNotFound
	}
	return c, nil
}

func (repo *InMemoryCountryRepository) SaveOrUpdate(_ context.Context, c *Country) error {
	repo.Lock()
	defer repo.Unlock()
	repo.storage[c.name] = c
	return nil
}
