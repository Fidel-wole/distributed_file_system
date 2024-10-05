package p2p

import "errors"

// ErrInvalidHandshake is returned when the handshake is invalid,
// if the handshake between the local and remote node fails.
var ErrInvalidHandshake = errors.New("invalid handshake")

// HandshakeFunc is a function that is called when a new connection is established.
type HandshakeFunc func(any) error

// NOPHandshakeFunc is a no-op handshake function that does nothing.
func NOPHandshakeFunc(any) error { return nil }