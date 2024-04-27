package MyHashMap

import (
	"crypto"
	"encoding/json"
	"log"
	"math"
)

type mapNode[K, E any] struct {
	key   K
	value E
}

type MyHashMap[K, E any] struct {
	chains   []*mapNode[K, E]
	capacity int
	harsher  crypto.Hash
	size     int
}

func NewMyHashMap[K, E any](initialCap int) *MyHashMap[K, E] {
	return &MyHashMap[K, E]{
		chains:   make([]*mapNode[K, E], 0),
		capacity: initialCap,
		harsher:  crypto.SHA256,
		size:     0,
	}
}

func (m *MyHashMap[K, E]) hash(el K) int {
	strEl, err := json.Marshal(el)
	if err != nil {
		log.Panicf("Hash, Marshal err: %v", err)
	}
	tempHarsher := m.harsher.New()
	tempHarsher.Write(strEl)
	salt := tempHarsher.Sum(nil)
	var intSalt int
	for i, letter := range salt {
		if intSalt > int(math.Pow10(7)) {
			break
		}
		intSalt += int(math.Pow10(i)) + int(letter)
	}
	intSalt %= m.capacity
	return intSalt
}

func (m *MyHashMap[K, E]) Put(key K, value E) {
	insertIndex := m.hash(key)
	m.chains[insertIndex] = &mapNode[K, E]{
		key:   key,
		value: value,
	}
}

func (m *MyHashMap[K, E]) Get(key K) E {
	searchedIndex := m.hash(key)
	return m.chains[searchedIndex].value
}

func (m *MyHashMap[K, E]) Delete(key K) {
	deletedIndex := m.hash(key)
	m.chains[deletedIndex] = nil
}
