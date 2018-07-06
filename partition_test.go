package ksuid

import (
	"testing"
)

func TestStringPartitioner(t *testing. T) {
	p1 := StringPartitioner("test")
	p2 := StringPartitioner("part")
	
	if v1 := p1(); v1 != 0x74657374 {
		t.Errorf("Expected partition 0x74657374, got %08x\n", v1)
		t.Fail()
	}
	
	if v2 := p2(); v2 != 0x70617274 {
		t.Errorf("Expected partition 0x70617274, got %08x\n", v2)
		t.Fail()
	}
}

func BenchmarkStringPartitioner(b *testing.B) {
	p := StringPartitioner("test")

	for i := 0; i < b.N; i++ {
		p()
	}
}

func TestNilPartitioner(t *testing.T) {
	if v := NilPartitioner(); v != 0 {
		t.Fail()
	}
}

func TestMacPartitioner(t *testing.T) {
	fn, err := MacPartitioner()

	if err != nil || fn == nil {
		t.Error(err)
		t.FailNow()
	}

	if fn() == 0 {
		t.Error("MacPartitioner()() must be >0")
		t.Fail()
	}
}
func BenchmarkMacPartitioner(b *testing.B) {
	p, err := MacPartitioner()

	if err != nil {
		b.FailNow()
	}

	for i := 0; i < b.N; i++ {
		p()
	}
}