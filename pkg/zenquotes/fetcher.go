package zenquotes

import (
	"encoding/json"
	"net/http"

	serviceErrors "github.com/dolbik/pow/pkg/error"
)

const QuotesAddr = "https://zenquotes.io"

type Fetcher interface {
	Fetch() (*Quote, error)
}

type ZendQuotes struct {
	addr string
}

func NewFetcher(addr string) Fetcher {
	return &ZendQuotes{
		addr: addr,
	}
}

func (z *ZendQuotes) Fetch() (*Quote, error) {
	response, err := http.Get(z.addr + "/api/random")
	if err != nil {
		return nil, err
	}

	if response.StatusCode > 399 {
		return nil, serviceErrors.ErrZenQuotesWrongHTTPCode
	}

	var quotes []*json.RawMessage
	err = json.NewDecoder(response.Body).Decode(&quotes)
	if err != nil {
		return nil, err
	}

	for _, quote := range quotes {
		var q *Quote
		raw, err := quote.MarshalJSON()
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(raw, &q)
		if err != nil {
			return nil, err
		}

		return q, nil
	}

	return nil, nil
}
