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

func TestInTime(t *testing.T) {
	now := time.Date(2018, 5, 1, 0, 0, 0, 0, time.UTC)

	candidates := []struct {
		t      time.Time
		inTime bool
	}{
		{
			t:      time.Date(2018, 5, 1, 0, 0, -1, 0, time.UTC),
			inTime: true,
		},
		{
			t:      now,
			inTime: true,
		},
		{
			t:      time.Date(2018, 5, 1, 0, 0, 10, 0, time.UTC),
			inTime: true,
		},
		{
			t:      time.Date(2018, 5, 1, 0, 0, 11, 0, time.UTC),
			inTime: true,
		},
	}

	for i, v := range candidates {
		if e, g := v.inTime, InTime(now, v.t, 10*time.Second); e != g {
			t.Fatalf("%d : expected %t; got %t", i, e, g)
		}
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
