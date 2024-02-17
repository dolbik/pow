package processor

import (
	"testing"

	"github.com/dolbik/pow/pkg/hashcash"
	"github.com/stretchr/testify/assert"
)

func TestChallengeProcessor_Process(t *testing.T) {
	prc := NewChallengeProcessor()
	hc := hashcash.NewHashCash("test", 0)
	body, _ := hc.Marshal()
	msg := &Message{
		Header: ChallengeResponse,
		Body:   body,
	}

	result, err := prc.Process(msg)

	assert.NoError(t, err)
	assert.Equal(t, ResourceRequest, result.Header)
	assert.Equal(t, body, result.Body)
}
