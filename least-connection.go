package main

import (
	"sync"
)

// Reference used: https://levelup.gitconnected.com/implementation-of-thread-safe-dictionary-data-structure-in-golang-2bcb235fd9e4 

// #-----------------------------------------------------------------------#
// #                          Least Connections                            #
// #-----------------------------------------------------------------------#

//Global Variables
// Initialise the dictornary.
var serverPool = Dictionary{}

type IKey interface{}
type IValue interface{}

// Dictionary struct, makes the map into a dictionary.
type Dictionary struct {
	items map[IKey]IValue
	lock  sync.RWMutex
}

// Add adds a key to the dictionary.
func (dict *Dictionary) Add(key IKey, value IValue) {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	if dict.items == nil {
		dict.items = make(map[IKey]IValue)
	}
	dict.items[key] = value
}

// Get values and insert to "GetUrl" variable as a string.
func (dict *Dictionary) GetUrl() string {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	var temp = 1000000000
	var target = ""
	for i := range dict.items {
		if dict.items[i].(int) < temp {
			target = i.(string)
			temp = dict.items[i].(int)
		}
	}
	return target
}

// AddConn adds a value to a key in the dictionary.
func (dict *Dictionary) AddConn(key IKey) {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	if dict.items == nil {
		dict.items = make(map[IKey]IValue)
	}
	dict.items[key] = dict.items[key].(int) + 1
}

// RemoveConn removes a value from a key in the dictionary.
func (dict *Dictionary) RemoveConn(key IKey) {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	if dict.items == nil {
		dict.items = make(map[IKey]IValue)
	}
	dict.items[key] = dict.items[key].(int) - 1
}

// GetValues returns a slice of all the values.
func (dict *Dictionary) GetValues() []IValue {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	values := []IValue{}
	for i := range dict.items {
		values = append(values, dict.items[i])
	}
	return values
}

// GetKeys returns a slice of all the keys.
func (dict *Dictionary) GetKeys() []IKey {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	keys := []IKey{}
	for i := range dict.items {
		keys = append(keys, i)
	}
	return keys
}
