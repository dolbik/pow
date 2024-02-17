package zenquotes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuote_String(t *testing.T) {
	q := &Quote{
		Q: "quote",
		A: "author",
		H: "",
	}

	assert.Equal(t, "quote (c)author", q.String())
}
