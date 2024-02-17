package message

import (
	"bytes"
	"errors"

	"github.com/dolbik/pow/pkg/message/processor"
	"github.com/dolbik/pow/pkg/zenquotes"
)

var (
	errInvalidHeader = errors.New("header in the message is not valid")
)

func GetMessageProcessor(msg *processor.Message) (processor.Processor, error) {

	if bytes.Equal(msg.Header, processor.ServiceRequest) {
		return processor.NewServiceProcessor(), nil
	}

	if bytes.Equal(msg.Header, processor.ChallengeResponse) {
		return processor.NewChallengeProcessor(), nil
	}

	if bytes.Equal(msg.Header, processor.ResourceRequest) {
		return processor.NewResourceProcessor(zenquotes.NewFetcher(zenquotes.QuotesAddr)), nil
	}

	return nil, errInvalidHeader
}
