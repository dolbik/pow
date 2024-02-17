package processor

import (
	"github.com/dolbik/pow/pkg/hashcash"
	"github.com/dolbik/pow/pkg/zenquotes"
)

type ResourceProcessor struct {
	quotesFetcher zenquotes.Fetcher
}

func NewResourceProcessor(quotesFetcher zenquotes.Fetcher) Processor {
	return &ResourceProcessor{
		quotesFetcher: quotesFetcher,
	}
}

func (rp *ResourceProcessor) Process(msg *Message) (*Message, error) {
	hc := new(hashcash.HashCash)
	err := hc.Unmarshal(msg.Body)

	if err != nil {
		return nil, err
	}

	err = hc.Verify(msg.Resource)
	if err != nil {
		return nil, err
	}

	quote, err := rp.quotesFetcher.Fetch()
	if err != nil {
		return nil, err
	}

	return &Message{
		Header: GrantResponse,
		Body:   []byte(quote.String()),
	}, nil
}
