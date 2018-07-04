package ksuid

import (
	"encoding/binary"
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
