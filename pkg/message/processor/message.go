package processor

import (
	"fmt"
	"strings"

	serviceErrors "github.com/dolbik/pow/pkg/error"
)

var (
	ServiceRequest    = []byte("service_request")
	ChallengeResponse = []byte("challenge_response")
	ResourceRequest   = []byte("resource_request")
	GrantResponse     = []byte("grant_response")
)

type Processor interface {
	Process(msg *Message) (*Message, error)
}

type Message struct {
	Header   []byte
	Body     []byte
	Resource string
}

func Parse(msg []byte) (*Message, error) {
	parts := strings.Split(strings.TrimSpace(string(msg)), "\n")
	if len(parts) < 1 || len(parts) > 2 {
		return nil, serviceErrors.ErrInvalidMessageFormat
	}

	body := ""
	if len(parts) == 2 {
		body = parts[1]
	}

	return &Message{
		Header: []byte(parts[0]),
		Body:   []byte(body),
	}, nil
}

func (m Message) Marshal() []byte {
	return []byte(fmt.Sprintf("%s\n%s", m.Header, m.Body))
}
