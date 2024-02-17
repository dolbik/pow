package cache

import (
	"sync"
)

var storage = map[int]bool{}
var mutex = &sync.RWMutex{}

func Add(key int) bool {
	if !Exists(key) {
		mutex.Lock()
		storage[key] = true
		mutex.Unlock()
		return true
	}
	return false
}

func Exists(key int) bool {
	mutex.RLock()
	defer mutex.RUnlock()
	_, ok := storage[key]
	return ok
}

func Delete(key int) {
	if Exists(key) {
		mutex.Lock()
		delete(storage, key)
		mutex.Unlock()
	}
}
