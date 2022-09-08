package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

// func TestHello(t *testing.T) {

// 	assertCorrectMessage := func(t testing.TB, got, want string) {
// 		t.Helper()
// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	}

// 	t.Run("saying hello to people", func(t *testing.T) {
// 		got := Hello("Chris", "")
// 		want := "Hello, Chris"

// 		assertCorrectMessage(t, got, want)
// 	})

// 	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
// 		got := Hello("", "")
// 		want := "Hello, world"

// 		assertCorrectMessage(t, got, want)
// 	})

// 	t.Run("in Spanish", func(t *testing.T) {
// 		got := Hello("Elodie", "Spanish")
// 		want := "Hola, Elodie"

// 		assertCorrectMessage(t, got, want)
// 	})

// 	t.Run("in French", func(t *testing.T) {
// 		got := Hello("Bla", "French")
// 		want := "Bonjour, Bla"

// 		assertCorrectMessage(t, got, want)
// 	})
// 	t.Run("in Romanian", func(t *testing.T) {
// 		got := Hello("Abc", "Romanian")
// 		want := "Buna, Abc"

// 		assertCorrectMessage(t, got, want)
// 	})
// }

// func BenchmarkRepeat(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Repeat("a")
// 	}
// }

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("got %d want %d", spySleeper.Calls, 3)
	}

	t.Run("sleep before every print", func(t *testing.T) {
		spySleeperPrinter := &CountdownOperationsSpy{}
		Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeperPrinter.Calls) {
			t.Errorf("wanted calls %v, got %v", want, spySleeperPrinter.Calls)
		}
	})
}

func TestConfugurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
