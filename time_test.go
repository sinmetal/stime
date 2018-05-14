package stime

import (
	"testing"
	"time"
)

func TestNow1(t *testing.T) {
	if Now().IsZero() {
		t.Fatalf("time is zero")
	}
}

func TestAddDummyTime(t *testing.T) {
	t1 := time.Date(2018, 5, 1, 1, 0, 0, 0, time.UTC)
	t2 := time.Date(2018, 5, 1, 0, 0, 0, 0, time.UTC)
	AddDummyTime(t1, t2)
	if a, b := Now(), Now(); a.After(b) {
		t.Fatalf("%s is %s", a, b)
	}
}

func TestSetPermafrost(t *testing.T) {
	t1 := time.Date(2018, 5, 1, 1, 0, 0, 0, time.UTC)
	SetPermafrost(t1)
	if a, b := Now(), t1; a.Equal(b) == false {
		t.Fatalf("%s is not equal %s", a, b)
	}
}
