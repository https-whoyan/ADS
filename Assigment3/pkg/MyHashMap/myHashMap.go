package MyHashMap

import (
	"cmp"
	"crypto"
	"encoding/json"
	"log"
	"math"
)

type ord cmp.Ordered

type mapNode[K ord, E any] struct {
	key   K
	value E
}

type MyHashMap[K ord, E any] struct {
	chains   []*mapNode[K, E]
	capacity int
	hasher   crypto.Hash
	size     int
}

func NewMyHashMap[K ord, E any](initialCap int) *MyHashMap[K, E] {
	return &MyHashMap[K, E]{
		chains:   make([]*mapNode[K, E], initialCap),
		capacity: initialCap,
		hasher:   crypto.SHA256,
		size:     0,
	}
}

func (m *MyHashMap[K, E]) hash(el K) int {
	// GerStrHash
	strEl, err := json.Marshal(el)
	if err != nil {
		log.Panicf("Hash, Marshal err: %v", err)
	}
	// Get tempHasher for hashing key
	tempHasher := m.hasher.New()
	tempHasher.Write(strEl)
	salt := tempHasher.Sum(nil)
	// ok, salt gotten, get the intSalt
	var intSalt int
	// I used ASCII ids
	// I increment 10 to the degree of the character's index order
	// multiply by ASCII id of letter
	for i, letter := range salt {
		// I assume that more 5 * m.capacity will give evenly different hashes.
		if intSalt > 5*m.capacity {
			break
		}
		intSalt += int(math.Pow10(i)) * int(letter)
	}
	// limit the hash modulus to the size of the hash map
	intSalt %= m.capacity
	return intSalt
}

func (m *MyHashMap[K, E]) Put(key K, value E) (haveACollision bool, oldKey K, oldVal E) {
	insertIndex := m.hash(key)
	if m.chains[insertIndex] == nil {
		m.size++
	} else {
		// Have a collision?
		// Stand haveACollision = true
		// Return old HashNode values
		haveACollision = true
		oldKey = m.chains[insertIndex].key
		oldVal = m.chains[insertIndex].value
	}
	m.chains[insertIndex] = &mapNode[K, E]{
		key:   key,
		value: value,
	}
	return
}

func (m *MyHashMap[K, E]) Get(key K) (value E, contains bool) {
	contains = true
	searchedIndex := m.hash(key)
	if m.chains[searchedIndex] == nil || m.chains[searchedIndex].key != key {
		contains = false
		return
	}
	value = m.chains[searchedIndex].value
	return
}

func (m *MyHashMap[K, E]) Remove(key K) {
	deletedIndex := m.hash(key)
	if m.chains[deletedIndex] != nil && m.chains[deletedIndex].key == key {
		m.size--
		m.chains[deletedIndex] = nil
	}
}

func (m *MyHashMap[K, E]) Contains(key K) bool {
	searchedIndex := m.hash(key)
	if m.chains[searchedIndex] == nil {
		return false
	}
	return m.chains[searchedIndex].key == key
}

func (m *MyHashMap[K, E]) IsEmpty() bool {
	return m.size == 0
}

func (m *MyHashMap[K, E]) Keys() []K {
	var keys []K
	for _, chain := range m.chains {
		if chain != nil {
			keys = append(keys, chain.key)
		}
	}
	return keys
}

func (m *MyHashMap[K, E]) Values() []E {
	var values []E
	for _, chain := range m.chains {
		if chain != nil {
			values = append(values, chain.value)
		}
	}
	return values
}

func (m *MyHashMap[K, E]) Clear() {
	m.size = 0
	m.chains = make([]*mapNode[K, E], m.capacity)
}

func (m *MyHashMap[K, E]) Size() int {
	return m.size
}
