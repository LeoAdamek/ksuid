package ksuid

import (
	"time"
	)

// Epoch gets the Epoch time from when timestamps are generated
// Defaults to 2000-01-01T00:00:00Z but can be overriden if required.
var Epoch = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).Unix()

// KSUID is the main representation for a K-Sortable Unique Identifier.
type KSUID struct {
	T time.Time
	Seq uint32
	Partition uint32
}

// Epoch gets the time the KSUID was generated based on the Epoch value.
func (k KSUID) Epoch() uint32 {
	return uint32(k.T.Unix() - Epoch)
}

// Binary is an alias for EncodeBinary(k)
func (k KSUID) Binary() []byte {
	return EncodeBinary(k)
}

// String returns the KSUID as a string using EncodeHex
func (k KSUID) String() string {
	return string(EncodeHex(k))
}