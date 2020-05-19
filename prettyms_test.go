package prettyms

import (
	"testing"

	parsems "github.com/fernandoporazzi/parse-ms"
)

func TestHumanize(t *testing.T) {
	got := Humanize(float64(1337000000))
	want := parsems.ParsedMilliseconds{
		Days:         15,
		Hours:        11,
		Minutes:      23,
		Seconds:      20,
		Milliseconds: 0,
		Microseconds: 0,
		Nanoseconds:  0,
	}

	if got != want {
		t.Errorf("Expected %v to be equal %v", got, want)
	}
}
