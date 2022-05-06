package inmemory

import (
	"fmt"
	"sync"
)

type datastore struct {
	data map[string]string
	*sync.RWMutex
}

type repository struct {
	datastore *datastore
}

func NewRepository() Repository {
	return &repository{
		datastore: &datastore{
			data:    make(map[string]string),
			RWMutex: &sync.RWMutex{},
		},
	}
}

func (r *repository) Get(key string) (*InMemory, error) {
	r.datastore.RLock()
	defer r.datastore.RUnlock()

	value, ok := r.datastore.data[key]
	if !ok {
		return nil, fmt.Errorf("key %s not found", key)
	}

	return &InMemory{
		Key:   key,
		Value: value,
	}, nil
}

func (r *repository) Set(key, value string) (*InMemory, error) {
	r.datastore.Lock()
	defer r.datastore.Unlock()
	_, ok := r.datastore.data[key]
	if ok {
		return nil, fmt.Errorf("key %s already exist", key)
	}
	r.datastore.data[key] = value

	return &InMemory{
		Key:   key,
		Value: value,
	}, nil
}
