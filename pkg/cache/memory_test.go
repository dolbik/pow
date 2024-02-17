package cache

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		key := i
		go func() {
			defer wg.Done()
			Add(key)
		}()
	}
	wg.Wait()

	for i := 0; i < 100; i++ {
		assert.True(t, Exists(i))
	}
}

func TestDelete(t *testing.T) {
	key := 12345
	Add(key)
	assert.True(t, Exists(key))
	Delete(key)
	assert.False(t, Exists(key))
}
