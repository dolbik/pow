package connection

import (
	"net"
	"testing"

	serviceErrors "github.com/dolbik/pow/pkg/error"
	"github.com/stretchr/testify/assert"
)

func TestIOReader_Read(t *testing.T) {
	r, w := net.Pipe()

	go func() {
		writer := NewIOWriter(w)
		_ = writer.Write([]byte("Hello, world!"))
		_ = w.Close()
	}()

	reader := NewIOReader(r)

	body, err := reader.Read()

	assert.NoError(t, err)
	assert.Equal(t, "Hello, world!", string(body))

	_ = r.Close()
}

func TestIOReader_Read_MessageHasNoHeader(t *testing.T) {
	r, w := net.Pipe()

	go func() {
		_, _ = w.Write([]byte("Hello, world!"))
		_ = w.Close()
	}()

	reader := NewIOReader(r)

	body, err := reader.Read()

	assert.Equal(t, serviceErrors.ErrMessageTooLong, err)
	assert.Nil(t, body)

	_ = r.Close()
}
