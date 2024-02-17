package processor

import (
	"testing"

	serviceErrors "github.com/dolbik/pow/pkg/error"
	"github.com/stretchr/testify/assert"
)

func TestParse_FormatIsInvalid(t *testing.T) {
	msg := []byte("a\nb\nc")
	_, err := Parse(msg)
	assert.Equal(t, serviceErrors.ErrInvalidMessageFormat, err)
}

func TestMessage_Marshal(t *testing.T) {
	msg := &Message{
		Header: []byte("header"),
		Body:   []byte("body"),
	}

	assert.Equal(t, []byte("header\nbody"), msg.Marshal())
}
