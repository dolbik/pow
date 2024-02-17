package processor

import (
	"github.com/dolbik/pow/pkg/hashcash"
)

type ServiceProcessor struct {
}

func NewServiceProcessor() Processor {
	return &ServiceProcessor{}
}

func (sp *ServiceProcessor) Process(msg *Message) (*Message, error) {
	hc := hashcash.NewHashCash(msg.Resource, hashcash.WantZeros)
	marshaled, err := hc.Marshal()
	if err != nil {
		return nil, err
	}

	return &Message{
		Header: ChallengeResponse,
		Body:   marshaled,
	}, nil
}
