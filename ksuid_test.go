package ksuid

import (
	"testing"
	"time"
)

var tst = KSUID{
T: time.Date(2018, 7, 4, 19, 51, 23, 0, time.UTC),
Seq: 0xC0FFEE,
Partition: 0xDEADBEEF,
}

func TestKSUID_String(t *testing.T) {

}
