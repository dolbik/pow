package message

import (
	"testing"

	process "github.com/dolbik/pow/pkg/message/processor"
	"github.com/stretchr/testify/assert"
)

func TestGetMessageProcessor_HeaderIsUnknown(t *testing.T) {
	_, err := GetMessageProcessor(&process.Message{Header: []byte("header")})
	assert.Equal(t, errInvalidHeader, err)
}

func TestGetMessageProcessor_ServiceRequest(t *testing.T) {
	processor, err := GetMessageProcessor(&process.Message{Header: process.ServiceRequest})
	assert.NoError(t, err)
	assert.IsType(t, new(process.ServiceProcessor), processor)
}

func TestGetMessageProcessor_ResourceRequest(t *testing.T) {
	processor, err := GetMessageProcessor(&process.Message{Header: process.ResourceRequest})
	assert.NoError(t, err)
	assert.IsType(t, new(process.ResourceProcessor), processor)
}

func TestGetMessageProcessor_ChallengeResponse(t *testing.T) {
	processor, err := GetMessageProcessor(&process.Message{Header: process.ChallengeResponse})
	assert.NoError(t, err)
	assert.IsType(t, new(process.ChallengeProcessor), processor)
}
