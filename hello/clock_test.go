package main

import (
	"testing"
	"time"

	"github.com/gypsydave5/learn-go-with-tests/math/v1/clockface"
)

func TestClockSecondHandAtMidnight(t *testing.T) {
	t.Logf("Test Convert Roman To Arabic Numerals")
	t.Logf("\t%s", t.Name())

	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(tm)

	if got == want {
		t.Logf("\t%s\tGot %v", succed, got)
	} else {
		t.Errorf("\t%s\tGot %v, want %v", failed, got, want)
	}
}
