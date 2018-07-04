package ksuid

import (
	"time"
	"sync/atomic"
)

// A Generator rep
type Generator interface {
	Next() KSUID
}

// AsyncGenerator optimizes for reduced allocations and comparisons by a background timer to
// Reset the sequence and change the timestamp each second.
// Note: This generator must be started with Run()
type AsyncGenerator struct {
	// The function used to partition KSUIDs
	Partitioner PartitionFunc
	clk *time.Ticker
	ts time.Time
	seq uint32
}

// StandardGenerator runs without goroutines or background tickers.
type StandardGenerator struct {
	Partitioner PartitionFunc
	ts time.Time
	seq uint32
}

var defaultGenerator = StandardGenerator{
	Partitioner: NilPartitioner,
}

// Next will get the next KSUID from the package-wide default generator.
func Next() KSUID {
	return defaultGenerator.Next()
}

// SetPartitioner sets the PartitionFunc for the package-wide default generator.
func SetPartitioner(p PartitionFunc) {
	defaultGenerator.Partitioner = p
}


func (g *StandardGenerator) Next() KSUID {
	t := time.Now().Truncate(time.Second)
	
	if t.After(g.ts) {
		atomic.StoreUint32(&g.seq, 0)
		g.ts = t
	}
	
	seq := atomic.LoadUint32(&g.seq)
	atomic.AddUint32(&g.seq, 1)
	
	return KSUID{
		T: t,
		Seq: seq,
		Partition: g.Partitioner(),
	}
}

func NewAsyncGenerator() *AsyncGenerator {
	return &AsyncGenerator{
		clk: time.NewTicker(time.Second),
		Partitioner: NilPartitioner,
	}
}

// Run will start the generator's background sequence.
func (g *AsyncGenerator) Run() {
	for {
		t := <-g.clk.C
		
		g.ts = t.Truncate(time.Second)
		atomic.StoreUint32(&g.seq, 0)
	}
}

// Next gets the next KSUID from the generator
func (g *AsyncGenerator) Next() KSUID {
	s := atomic.LoadUint32(&g.seq)
	
	atomic.AddUint32(&g.seq, 1)
	
	return KSUID{
		Partition: g.Partitioner(),
		Seq: s,
		T: g.ts,
	}
}
