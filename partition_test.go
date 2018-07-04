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

func TestNilPartitioner(t *testing.T) {
	if v := NilPartitioner(); v != 0 {
		t.Fail()
	}
}
