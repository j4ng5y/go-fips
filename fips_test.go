package fips

import (
	"testing"
)

func TestFIPS_FromString_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("NewFIPS().FromString(\"a\") did not panic")
		}
	}()

	_ = NewFIPS().FromString("a")
}

func TestFIPS_String(t *testing.T) {
	fipsZero := NewFIPS()
	if fipsZero != FIPS(0) {
		t.Logf("expected uint32:0, got uint32:%d", fipsZero)
		t.Fail()
	}

	fipsStr := NewFIPS().FromString("12345")
	if fipsStr != FIPS(uint32(12345)) {
		t.Logf("expected uint32:12345, got uint32:%d", fipsStr)
		t.Fail()
	}
	if fipsStr.String() != "12345" {
		t.Logf("excpected str:12345, got str:%s", fipsStr.String())
		t.Fail()
	}

	fipsStr = NewFIPS().FromInt(1234)
	if fipsStr.String() != "01234" {
		t.Logf("expected str:01234, got str:%s", fipsStr.String())
	}

	fipsStr = NewFIPS().FromInt8(1)
	if fipsStr.String() != "00001" {
		t.Logf("expected str:00001, got str:%s", fipsStr.String())
	}

	fipsStr = NewFIPS().FromInt16(1234)
	if fipsStr.String() != "01234" {
		t.Logf("expected str:01234, got str:%s", fipsStr.String())
	}

	fipsStr = NewFIPS().FromInt32(1234)
	if fipsStr.String() != "01234" {
		t.Logf("expected str:01234, got str:%s", fipsStr.String())
	}

	fipsStr = NewFIPS().FromInt64(1234)
	if fipsStr.String() != "01234" {
		t.Logf("expected str:01234, got str:%s", fipsStr.String())
	}

	fipsStr = NewFIPS().FromUint(1234)
	if fipsStr.String() != "01234" {
		t.Logf("expected str:01234, got str:%s", fipsStr.String())
	}

	fipsStr = NewFIPS().FromUint8(1)
	if fipsStr.String() != "00001" {
		t.Logf("expected str:00001, got str:%s", fipsStr.String())
	}

	fipsStr = NewFIPS().FromUint16(1234)
	if fipsStr.String() != "01234" {
		t.Logf("expected str:01234, got str:%s", fipsStr.String())
	}

	fipsStr = NewFIPS().FromUint32(1234)
	if fipsStr.String() != "01234" {
		t.Logf("expected str:01234, got str:%s", fipsStr.String())
	}

	fipsStr = NewFIPS().FromUint64(1234)
	if fipsStr.String() != "01234" {
		t.Logf("expected str:01234, got str:%s", fipsStr.String())
	}
}