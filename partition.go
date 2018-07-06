package ksuid

import (
	"encoding/binary"
	"net"
)

// PartitionFunc is a func which returns a partition for the
type PartitionFunc func() uint32

// StringPartitioner uses the first four bytes of the given string to partition.
func StringPartitioner(v string) PartitionFunc {
	val := binary.BigEndian.Uint32([]byte(v))
	
	return func() uint32 {
		return val
	}
}

// NilPartitioner places everything into partition zero, disabling partitioning
func NilPartitioner() uint32 {
	return 0
}

// MacPartitioner places everything into a partition based on the last 4 bytes of the primary network mac address
// It's a wrapped in a func which may error if
func MacPartitioner() (PartitionFunc, error) {

	interfaces, err := net.Interfaces()

	if err != nil { return nil, err }

	for _, i := range interfaces {
		if len(i.HardwareAddr) > 0 {
			a := i.HardwareAddr[2:]

			v := binary.BigEndian.Uint32(a)

			return func() uint32 {
				return v
			}, nil
		}
	}

	return nil, nil
}
