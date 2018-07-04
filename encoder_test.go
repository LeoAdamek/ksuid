package ksuid

import (
	"testing"
	)

func TestEncodeBinary(t *testing.T) {

	
	v := EncodeBinary(tst)
	expected := []byte{0x22, 0xcf, 0xe1, 0xbb, 0x00, 0xc0, 0xff, 0xee, 0xde, 0xad, 0xbe, 0xef}
	
	if len(v) != 12 {
		t.Errorf("Expected 12 bytes, got %d\n", len(v))
		t.FailNow()
	}
	
	for i := 0; i < len(expected); i++ {
		if v[i] != expected[i] {
			t.Errorf("Expected byte %d = 0x%02X got 0x%02X", i,  expected[i], v[i])
			
			t.Errorf("V: %+#v\n", v)
			t.Errorf("E: %+#v\n", expected)
			t.Fail()
		}
	}
	
}

func TestDecodeBinary(t *testing.T) {
	r := []byte{0x22, 0xcf, 0xe1, 0xbb, 0x00, 0xc0, 0xff, 0xee, 0xde, 0xad, 0xbe, 0xef}
	
	k, _ := DecodeBinary(r)
	
	if k.Seq != 0xc0ffee {
		t.Errorf("Expected Partition 0xc0ffee, got 0x%08X\n", k.Seq)
		t.Fail()
	}
}

func TestEncodeHex(t *testing.T) {
	v := EncodeHex(tst)
	
	if strLen := len(v); strLen != 24 {
		t.Errorf("Expected 24 chars of data, got %d\n", v)
		t.FailNow()
	}
	
	if string(v) != "22cfe1bb00c0ffeedeadbeef" {
		t.Errorf("Expected '22cfe1bb00c0ffeedeadbeef' got '%s'\n", string(v))
		t.FailNow()
	}
}

func TestDecodeHex(t *testing.T) {
	r := "22cfe1bb00c0ffeedeadbeef"
	
	k, err := DecodeHex([]byte(r))
	
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	
	if k.Seq != 0xC0FFEE {
		t.Errorf("Expected Partition 0xc0ffee, got 0x%08X\n", k.Seq)
		t.Fail()
	}
}

func BenchmarkEncodeBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncodeBinary(tst)
	}
}

func BenchmarkDecodeBinary(b *testing.B) {
	r := []byte{0x22, 0xcf, 0xe1, 0xbb, 0x00, 0xc0, 0xff, 0xee, 0xde, 0xad, 0xbe, 0xef}
	
	for i := 0; i < b.N; i++ {
		DecodeBinary(r)
	}
}

func BenchmarkDecodeHex(b *testing.B) {
	r := []byte("22cfe1bb00c0ffeedeadbeef")
	
	for i := 0; i < b.N; i++ {
		DecodeHex(r)
	}
}

func BenchmarkEncodeHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncodeHex(tst)
	}
}
