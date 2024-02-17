package hashcash

import (
	"math/rand"
	"testing"
	"time"

	"github.com/dolbik/pow/pkg/cache"
	serviceErrors "github.com/dolbik/pow/pkg/error"
	"github.com/stretchr/testify/assert"
)

func TestNewHashCash(t *testing.T) {
	result := new(HashCash)
	tm := time.Date(1979, 8, 2, 12, 34, 0, 0, time.UTC)
	hashStructure := &HashCash{
		ver:      1,
		bits:     2,
		date:     tm,
		resource: "test",
		random:   100,
		counter:  0,
	}

	marshaled, _ := hashStructure.Marshal()
	err := result.Unmarshal(marshaled)
	assert.NoError(t, err)
	assert.Equal(t, hashStructure, result)
}

func TestHashCash_Compute(t *testing.T) {
	tm := time.Date(1979, 8, 2, 12, 34, 0, 0, time.UTC)
	hashStructure := &HashCash{
		ver:      1,
		bits:     5,
		date:     tm,
		resource: "test",
		random:   100,
		counter:  0,
	}
	expected := &HashCash{
		ver:      1,
		bits:     5,
		date:     tm,
		resource: "test",
		random:   100,
		counter:  585697,
	}
	h, err := hashStructure.Compute()
	assert.NoError(t, err)
	assert.Equal(t, expected, h)
}

func TestHashCash_Compute_HugeBits(t *testing.T) {
	hashStructure := NewHashCash("test", 100)
	h, err := hashStructure.Compute()
	assert.Equal(t, serviceErrors.ErrZeroBytesCountOverflow, err)
	assert.Nil(t, h)
}

func TestHashCash_Verify(t *testing.T) {
	hashStructure := NewHashCash("test", 2)
	computed, _ := hashStructure.Compute()
	err := computed.Verify("test")
	assert.NoError(t, err)
}

func TestHashCash_Verify_ResourceMisMatch(t *testing.T) {
	hashStructure := NewHashCash("test", 2)
	computed, _ := hashStructure.Compute()
	err := computed.Verify("test1")
	assert.Equal(t, serviceErrors.ErrResourceMismatch, err)
}

func TestHashCash_Verify_HashNotAcceptable(t *testing.T) {
	hashStructure := NewHashCash("test", 2)
	computed, _ := hashStructure.Compute()
	computed.counter = 1
	err := computed.Verify("test")
	assert.Equal(t, serviceErrors.ErrHashNotAcceptable, err)
}

func TestHashCash_Verify_HashExpired(t *testing.T) {
	rnd := rand.Int()
	cache.Add(rnd)
	hashStructure := &HashCash{
		ver:      1,
		bits:     2,
		date:     time.Now().Add(-11 * time.Minute),
		resource: "test",
		random:   rnd,
		counter:  0,
	}
	computed, _ := hashStructure.Compute()
	err := computed.Verify("test")
	assert.Equal(t, serviceErrors.ErrHashExpired, err)
}

func TestHashCash_Verify_RandomUnknown(t *testing.T) {
	hashStructure := &HashCash{
		ver:      1,
		bits:     2,
		date:     time.Now(),
		resource: "test",
		random:   rand.Int(),
		counter:  0,
	}
	computed, _ := hashStructure.Compute()
	err := computed.Verify("test")
	assert.Equal(t, serviceErrors.ErrHashDoesNotExist, err)
}
