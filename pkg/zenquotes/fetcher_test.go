package zenquotes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	serviceErrors "github.com/dolbik/pow/pkg/error"
	"github.com/stretchr/testify/assert"
)

func TestZendQuotes_Fetch(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/api/random", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, `[ {"q":"Even in the grave, all is not lost.","a":"Edgar Allan Poe","h":"<blockquote>&ldquo;Even in the grave, all is not lost.&rdquo; &mdash; <footer>Edgar Allan Poe</footer></blockquote>"} ]`)
	})

	expectedQuote := &Quote{
		Q: "Even in the grave, all is not lost.",
		A: "Edgar Allan Poe",
		H: "<blockquote>&ldquo;Even in the grave, all is not lost.&rdquo; &mdash; <footer>Edgar Allan Poe</footer></blockquote>",
	}
	fetcher := NewFetcher(server.URL)
	quote, err := fetcher.Fetch()
	assert.NoError(t, err)
	assert.Equal(t, expectedQuote, quote)
}

func TestZendQuotes_HTTPError(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/api/random", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	})
	fetcher := NewFetcher(server.URL)
	quote, err := fetcher.Fetch()
	assert.Equal(t, serviceErrors.ErrZenQuotesWrongHTTPCode, err)
	assert.Nil(t, quote)
}
