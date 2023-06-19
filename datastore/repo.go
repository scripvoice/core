package datastore

import (
	"errors"
	"sync"
)

type Irepo interface {
	SetStore(store *SqlStore) Irepo
}

func GetRepo[T Irepo](s *SqlStore) (T, error) {
	t := new(T)
	var r Irepo = (*t).SetStore(s)
	v, ok := r.(T)
	if !ok {
		return *t, errors.New("unable to get base repo")
	}

	return v, nil
}

type BaseRepo[T any] struct {
	Store     *SqlStore
	storeInit sync.Once
}

func (r *BaseRepo[T]) SetStore(s *SqlStore) Irepo {
	r.storeInit.Do(func() {
		r.Store = s
	})
	return r
}

func (r *BaseRepo[T]) Save(entity T) {
	r.Store.db.Save(entity)
}
