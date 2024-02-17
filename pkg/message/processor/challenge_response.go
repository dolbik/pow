package processor

import (
	"github.com/dolbik/pow/pkg/hashcash"
)

type ChallengeProcessor struct {
}

func NewChallengeProcessor() Processor {
	return &ChallengeProcessor{}
}

func (rp *ChallengeProcessor) Process(msg *Message) (*Message, error) {
	hc := new(hashcash.HashCash)
	err := hc.Unmarshal(msg.Body)
	if err != nil {
		return nil, err
	}

	result, err := hc.Compute()
	if err != nil {
		return nil, err
	}

	marshaled, err := result.Marshal()
	if err != nil {
		return nil, err
	}

	return &Message{
		Header: ResourceRequest,
		Body:   marshaled,
	}, nil
}
