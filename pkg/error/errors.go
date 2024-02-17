package error

import "errors"

var (
	ErrHashNotAcceptable      = errors.New("the hash is not acceptable")
	ErrHashExpired            = errors.New("the hash expired")
	ErrResourceMismatch       = errors.New("resource mismatch")
	ErrHashDoesNotExist       = errors.New("the hash does not exist")
	ErrHashCompute            = errors.New("exceeded 1000000 iterations failed to find solution")
	ErrZeroBytesCountOverflow = errors.New("zero bytes count more than hash length")
	ErrUnmarshalHash          = errors.New("cannot unmarshal hash to struct")

	ErrMessageTooLong       = errors.New("message to long")
	ErrInvalidMessageFormat = errors.New("message format is not valid")

	ErrZenQuotesWrongHTTPCode = errors.New("ZendQuotes returned unexpected HTTP code")

	ErrServerAddress = errors.New("server host should be set")
	ErrServerPort    = errors.New("server port should be set")
)
