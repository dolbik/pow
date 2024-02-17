package hashcash

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/dolbik/pow/pkg/cache"
	serviceErrors "github.com/dolbik/pow/pkg/error"
)

const (
	iterations        = 1 << 32
	zero         byte = 48
	timeFormat        = "20060102150405"
	expireWindow      = 10 * time.Minute
	WantZeros         = 2
)

type Marshaler interface {
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
}

type HashCashier interface {
	Compute() (*HashCash, error)
	Verify() bool
}

type HashCash struct {
	ver      int
	bits     int
	date     time.Time
	resource string
	random   int
	counter  int
}

func NewHashCash(resource string, zeroBits int) *HashCash {
	rnd := rand.Int()
	for !cache.Add(rnd) {
		rnd = rand.Int()
	}
	return &HashCash{
		ver:      1,
		bits:     zeroBits,
		date:     time.Now().UTC(),
		resource: resource,
		random:   rnd,
		counter:  0,
	}
}

func (h *HashCash) Marshal() ([]byte, error) {
	var b bytes.Buffer
	_, _ = fmt.Fprintf(
		&b,
		"%d:%d:%s:%s::%s:%s",
		h.ver,
		h.bits,
		h.date.Format(timeFormat),
		base64.StdEncoding.EncodeToString([]byte(h.resource)),
		base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", h.random))),
		base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", h.counter))),
	)

	return b.Bytes(), nil
}

func (h *HashCash) Unmarshal(data []byte) error {
	var err error
	parts := strings.Split(string(data), ":")

	if len(parts) != 7 {
		return serviceErrors.ErrUnmarshalHash
	}

	h.ver, err = strconv.Atoi(parts[0])
	if err != nil {
		return err
	}

	h.bits, err = strconv.Atoi(parts[1])
	if err != nil {
		return err
	}

	h.date, err = time.Parse(timeFormat, parts[2])
	if err != nil {
		return err
	}

	str, err := base64.StdEncoding.DecodeString(parts[3])
	if err != nil {
		return err
	}
	h.resource = string(str)

	str, err = base64.StdEncoding.DecodeString(parts[5])
	if err != nil {
		return err
	}
	h.random, err = strconv.Atoi(string(str))
	if err != nil {
		return err
	}

	str, err = base64.StdEncoding.DecodeString(parts[6])
	if err != nil {
		return err
	}
	h.counter, err = strconv.Atoi(string(str))
	if err != nil {
		return err
	}

	return err
}

func (h *HashCash) Compute() (*HashCash, error) {
	var hash []byte
	for {
		plain, err := h.Marshal()
		if err != nil {
			return nil, err
		}

		hash, err = h.sha256(plain)
		if err != nil {
			return nil, err
		}

		if h.bits >= len(hash) {
			return nil, serviceErrors.ErrZeroBytesCountOverflow
		}

		if h.isAcceptable(hash, h.bits) {
			break
		}

		h.counter++
		if h.counter > iterations {
			return nil, serviceErrors.ErrHashCompute
		}
	}

	return h, nil
}

func (h *HashCash) Verify(resource string) error {
	plain, err := h.Marshal()
	if err != nil {
		return err
	}

	hash, err := h.sha256(plain)

	if h.resource != resource {
		return serviceErrors.ErrResourceMismatch
	}
	if h.date.Add(expireWindow).Before(time.Now()) {
		return serviceErrors.ErrHashExpired
	}

	if !h.isAcceptable(hash, h.bits) {
		return serviceErrors.ErrHashNotAcceptable
	}

	if !cache.Exists(h.random) {
		return serviceErrors.ErrHashDoesNotExist
	}

	return nil
}

func (h *HashCash) sha256(plaintext []byte) ([]byte, error) {
	hash := sha256.New()
	_, err := hash.Write(plaintext)
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("%x", hash.Sum(nil))), nil
}

func (h *HashCash) isAcceptable(hash []byte, bits int) bool {
	for _, val := range hash[:bits] {
		if val != zero {
			return false
		}
	}
	return true
}
