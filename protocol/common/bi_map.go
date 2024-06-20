package common

import (
	"sync"
)

type BiMap[K1 comparable, K2 comparable] struct {
	key1ToKey2 map[K1]K2
	key2ToKey1 map[K2]K1
	mutex      sync.RWMutex
}

func NewBiMap[K1 comparable, K2 comparable]() *BiMap[K1, K2] {
	return &BiMap[K1, K2]{
		key1ToKey2: make(map[K1]K2),
		key2ToKey1: make(map[K2]K1),
	}
}

func (b *BiMap[K1, K2]) Insert(key1 K1, key2 K2) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.key1ToKey2[key1] = key2
	b.key2ToKey1[key2] = key1
}

func (b *BiMap[K1, K2]) RetrieveByKey1(key1 K1) (K2, bool) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	key2, exists := b.key1ToKey2[key1]
	return key2, exists
}

func (b *BiMap[K1, K2]) RetrieveByKey2(key2 K2) (K1, bool) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	key1, exists := b.key2ToKey1[key2]
	return key1, exists
}

func (b *BiMap[K1, K2]) RemoveByKey1(key1 K1) bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	key2, exists := b.key1ToKey2[key1]
	if exists {
		delete(b.key1ToKey2, key1)
		delete(b.key2ToKey1, key2)
	}
	return exists
}

func (b *BiMap[K1, K2]) RemoveByKey2(key2 K2) bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	key1, exists := b.key2ToKey1[key2]
	if exists {
		delete(b.key2ToKey1, key2)
		delete(b.key1ToKey2, key1)
	}
	return exists
}
