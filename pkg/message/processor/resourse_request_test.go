package processor

import (
	"testing"

	"github.com/dolbik/pow/pkg/hashcash"
	"github.com/dolbik/pow/pkg/zenquotes"
	"github.com/stretchr/testify/assert"
)

func TestResourceProcessor_Process(t *testing.T) {
	fFetcher := new(fakeFetcher)
	process := NewResourceProcessor(fFetcher)
	hc := hashcash.NewHashCash("test", 2)
	computed, _ := hc.Compute()
	body, _ := computed.Marshal()
	msg := &Message{
		Header:   ResourceRequest,
		Body:     body,
		Resource: "test",
	}

	response, err := process.Process(msg)

	assert.NoError(t, err)
	assert.Equal(t, GrantResponse, response.Header)
	assert.Equal(t, []byte("quote (c)author"), response.Body)
}

type fakeFetcher struct{}

func (ff *fakeFetcher) Fetch() (*zenquotes.Quote, error) {
	return &zenquotes.Quote{
		Q: "quote",
		A: "author",
		H: "",
	}, nil
}
