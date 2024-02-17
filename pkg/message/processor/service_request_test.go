package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceProcessor_Process(t *testing.T) {
	sp := NewServiceProcessor()
	msg, err := sp.Process(&Message{
		Header: ServiceRequest,
		Body:   []byte(""),
	})
	assert.NoError(t, err)
	assert.Equal(t, ChallengeResponse, msg.Header)
}
