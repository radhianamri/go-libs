package errors

import (
	"errors"
	"io"
	"strings"
)

// to identify whether there were any netork connection issues during HTTP or RPC call
// such cases should be retriable since the network request has not been fully
// received by the remote server
// Note:
// Defined errors below are fully dependent on the library/go version used.
// Any new error type will require to be added later in future development
func IsNetwork(err error) bool {
	if err == nil {
		return false
	}

	// commonly found as underlying err from http transport during shut down race,
	// where the server closed the connection at the same time the client wrote
	// source: https://github.com/golang/go/blob/3fcbfb07a82c5332e6b50cddba333af6e6e3e488/src/net/http/transport.go#L883
	if errors.Is(err, io.EOF) {
		return true
	}

	// returned when a network descriptor is used after it has been closed
	// source: https://github.com/golang/go/blob/6d1c507bfc360ba72ca716bb7cb7bd9105a45af4/src/internal/poll/fd.go#L20-L31
	if strings.Contains(err.Error(), "use of closed network connection") {
		return true
	}

	// OS defined err returned during syscall when remote server returns a RST packet
	// during http handshake dropping the connection.
	// explanation: https://stackoverflow.com/a/1434506/8101467
	// source: https://github.com/golang/go/blob/6d1c507bfc360ba72ca716bb7cb7bd9105a45af4/src/syscall/tables_js.go#L205
	if strings.Contains(err.Error(), "connection reset by peer") {
		return true
	}

	// error defined by in old versions of lib https://github.com/go-micro/go-micro
	// returned when reusing a client connection that has already been closed in server
	// source: https://github.com/go-micro/go-micro/blob/a9be1288d2b63b8ab4f7b96d0b9b6b0bc3988ff4/transport/memory/memory.go#L73
	if strings.Contains(err.Error(), "server connection close") {
		return true
	}

	return false
}

func Contains(slice []error, err error) bool {
	for i := range slice {
		if errors.Is(slice[i], err) {
			return true
		}
	}
	return false
}

func Cause(err error) error {
	if err == nil {
		return nil
	}

	cause := err
	for {
		unwrapErr, ok := cause.(interface{ Unwrap() error })
		if !ok {
			break
		}
		cause = unwrapErr.Unwrap()
	}

	return cause
}
