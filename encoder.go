package ksuid

import (
		"bytes"
	"encoding/binary"
	"encoding/hex"
	"time"
	"errors"
)

// Encoder is a function which takes a KSUID and returns a sequence of bytes to represent it.
type Encoder func(KSUID) []byte

// Decoder is a function which takes a sequence of bytes and returns a KSUID
type Decoder func([]byte) (*KSUID, error)

// EncodeBinary encodes the KSUID as a sequence of bytes in big-endian order.
// Useful for efficient transport over a network to other machines.
// Also used by most other encoder functions.
func EncodeBinary(k KSUID) []byte {
	b := new(bytes.Buffer)
	
	binary.Write(b, binary.BigEndian, k.Epoch())
	binary.Write(b, binary.BigEndian, k.Seq)
	binary.Write(b, binary.BigEndian, k.Partition)
	
	return b.Bytes()
}

// EncodeHex encodes the KSUID to hex
func EncodeHex(k KSUID) []byte {
	b := EncodeBinary(k)
	
	r := make([]byte, 2*len(b))
	hex.Encode(r, b)
	
	return r
}

func DecodeBinary(b []byte) (*KSUID, error) {
	if len(b) != 12 {
		return nil, errors.New("encoded KSUID must be exactly 12 bytes")
	}
	
	t := binary.BigEndian.Uint32(b)
	s := binary.BigEndian.Uint32(b[4:])
	p := binary.BigEndian.Uint32(b[8:])
	
	k := &KSUID{
		T: time.Unix(Epoch + int64(t), 0),
		Seq: s,
		Partition: p,
	}
	
	return k, nil
}

func DecodeHex(b []byte) (*KSUID, error) {
	v, err := hex.DecodeString(string(b))
	
	if err != nil {
		return nil, err
	}
	
	return DecodeBinary(v)
}