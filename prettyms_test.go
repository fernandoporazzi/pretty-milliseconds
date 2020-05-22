package prettyms

import (
	"testing"
)

func TestHumanize(t *testing.T) {
	t.Run("Humanize function with default options", func(t *testing.T) {
		entries := []struct {
			input float64
			want  string
		}{
			{0, "0ms"},
			{0.1, "1ms"},
			{999, "999ms"},
			{1000, "1s"},
			{1400, "1.4s"},
			{2400, "2.4s"},
			{55000, "55s"},
			{67000, "1m 7s"},
			{300000, "5m"},
			{4020000, "1h 7m"},
			{43200000, "12h"},
			{144000000, "1d 16h"},
			{3596400000, "41d 15h"},
			{40176000000, "1y 100d"},
			{44863200000, "1y 154d 6h"},
			{119999, "1m 59.9s"},
			{120000, "2m"},
		}

		for _, entry := range entries {
			got := Humanize(entry.input)

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})

	t.Run("have a compact option", func(t *testing.T) {
		entries := []struct {
			input float64
			want  string
		}{
			{1004, "1s"},
			{3596400000, "41d"},
			{40176000000, "1y"},
			{44863200000, "1y"},
		}

		for _, entry := range entries {
			got := Humanize(entry.input, Options{
				Compact: true,
			})

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})

	t.Run("have a unitCount option", func(t *testing.T) {
		entries := []struct {
			input     float64
			unitCount int
			want      string
		}{
			{60000, 0, "1m"},
			{60000, 1, "1m"},
			{4020000, 1, "1h"},
			{4020000, 2, "1h 7m"},
			{44863200000, 1, "1y"},
			{44863200000, 2, "1y 154d"},
			{44863200000, 3, "1y 154d 6h"},
		}

		for _, entry := range entries {
			got := Humanize(entry.input, Options{
				UnitCount:     entry.unitCount,
				WithUnitCount: true,
			})

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})

	t.Run("have a SecondsDecimalDigits option", func(t *testing.T) {
		entries := []struct {
			input                float64
			secondsDecimalDigits int
			want                 string
		}{
			{999, 0, "999ms"},
			{1000, 0, "1s"},
			{1999, 0, "1s"},
			{2000, 0, "2s"},
			{33333, 0, "33s"},
			{33333, 2, "33.33s"},
			{33333, 3, "33.333s"},
		}

		for _, entry := range entries {
			got := Humanize(entry.input, Options{
				SecondsDecimalDigits:     entry.secondsDecimalDigits,
				WithSecondsDecimalDigits: true,
			})

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})

	t.Run("have a MillisecondsDecimalDigits option", func(t *testing.T) {
		entries := []struct {
			input                     float64
			millisecondsDecimalDigits int
			want                      string
		}{
			{33.333, 0, "33ms"},
			{33.3333, 4, "33.3333ms"},
			{12.123, 4, "12.123ms"},
		}

		for _, entry := range entries {
			got := Humanize(entry.input, Options{
				MillisecondsDecimalDigits: entry.millisecondsDecimalDigits,
			})

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})

	// Somehow this didn't work
	// Help wanted
	// t.Run("have a KeepDecimalsOnWholeSeconds option", func(t *testing.T) {
	// 	entries := []struct {
	// 		input                      float64
	// 		secondsDecimalDigits       int
	// 		keepDecimalsOnWholeSeconds bool
	// 		want                       string
	// 	}{
	// 		{33000, 3, true, "33ms"},
	// 	}

	// 	for _, entry := range entries {
	// 		got := Humanize(entry.input, Options{
	// 			SecondsDecimalDigits:       entry.secondsDecimalDigits,
	// 			KeepDecimalsOnWholeSeconds: entry.keepDecimalsOnWholeSeconds,
	// 			WithSecondsDecimalDigits:   true,
	// 		})

	// 		if got != entry.want {
	// 			t.Errorf("Expected %v to be equal %v", got, entry.want)
	// 		}
	// 	}
	// })

	t.Run("have a MillisecondsDecimalDigits option", func(t *testing.T) {
		entries := []struct {
			input float64
			want  string
		}{
			{0, "0 milliseconds"},
			{0.1, "1 millisecond"},
			{1, "1 millisecond"},
			{1000, "1 second"},
			{2400, "2.4 seconds"},
			{5000, "5 seconds"},
			{55000, "55 seconds"},
			{67000, "1 minute 7 seconds"},
			{300000, "5 minutes"},
			{4020000, "1 hour 7 minutes"},
			{43200000, "12 hours"},
			{144000000, "1 day 16 hours"},
			{3596400000, "41 days 15 hours"},
			{40176000000, "1 year 100 days"},
			{44863200000, "1 year 154 days 6 hours"},
			{649993232323, "20 years 223 days 1 hour 40 minutes 32.3 seconds"},
		}

		for _, entry := range entries {
			got := Humanize(entry.input, Options{
				Verbose: true,
			})

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})

	t.Run("have a SeparateMilliseconds option", func(t *testing.T) {
		entries := []struct {
			input                float64
			separateMilliseconds bool
			want                 string
		}{
			{1100, false, "1.1s"},
			{1100, true, "1s 100ms"},
		}

		for _, entry := range entries {
			got := Humanize(entry.input, Options{
				SeparateMilliseconds: entry.separateMilliseconds,
			})

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})

	t.Run("have a FormatSubMilliseconds option", func(t *testing.T) {
		entries := []struct {
			input float64
			want  string
		}{
			{0.4, "400µs"},
			{0.123571, "123µs 571ns"},
			{0.123456789, "123µs 456ns"},
			{3623433.123456, "1h 23s 433ms 123µs 456ns"},
		}

		for _, entry := range entries {
			got := Humanize(entry.input, Options{
				FormatSubMilliseconds: true,
			})

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})

	t.Run("have Verbose and Compact options", func(t *testing.T) {
		entries := []struct {
			input float64
			want  string
		}{
			{1000, "1 second"},
			{2400, "2 seconds"},
			{5000, "5 seconds"},
			{55000, "55 seconds"},
			{67000, "1 minute"},
			{300000, "5 minutes"},
			{4020000, "1 hour"},
			{43200000, "12 hours"},
			{144000000, "1 day"},
			{3596400000, "41 days"},
			{40176000000, "1 year"},
			{72360000000, "2 years"},
		}

		for _, entry := range entries {
			got := Humanize(entry.input, Options{
				Verbose: true,
				Compact: true,
			})

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})

	t.Run("have Verbose and UnitCount options", func(t *testing.T) {
		entries := []struct {
			input     float64
			unitCount int
			want      string
		}{
			{60000, 1, "1 minute"},
			{4020000, 1, "1 hour"},
			{4020000, 2, "1 hour 7 minutes"},
			{44863200000, 1, "1 year"},
			{44863200000, 2, "1 year 154 days"},
			{44863200000, 3, "1 year 154 days 6 hours"},
		}

		for _, entry := range entries {
			got := Humanize(entry.input, Options{
				Verbose:       true,
				UnitCount:     entry.unitCount,
				WithUnitCount: true,
			})

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})

	t.Run("have Verbose and FormatSubMilliseconds options", func(t *testing.T) {
		entries := []struct {
			input float64
			want  string
		}{
			{0.4, "400 microseconds"},
			{0.123571, "123 microseconds 571 nanoseconds"},
			{0.123456789, "123 microseconds 456 nanoseconds"},
			{0.001, "1 microsecond"},
		}

		for _, entry := range entries {
			got := Humanize(entry.input, Options{
				Verbose:               true,
				FormatSubMilliseconds: true,
			})

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})

	t.Run("compact option overrides unitCount option", func(t *testing.T) {
		entries := []struct {
			input     float64
			unitCount int
			want      string
		}{
			{44863200000, 1, "1 year"},
			{44863200000, 2, "1 year"},
			{44863200000, 3, "1 year"},
		}

		for _, entry := range entries {
			got := Humanize(entry.input, Options{
				Verbose:       true,
				Compact:       true,
				UnitCount:     entry.unitCount,
				WithUnitCount: true,
			})

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})

	t.Run("have SeparateMilliseconds and FormatSubMilliseconds option", func(t *testing.T) {
		entries := []struct {
			input float64
			want  string
		}{
			{1010.340067, "1s 10ms 340µs 67ns"},
			{60034.000005, "1m 34ms 5ns"},
		}

		for _, entry := range entries {
			got := Humanize(entry.input, Options{
				SeparateMilliseconds:  true,
				FormatSubMilliseconds: true,
			})

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})

	t.Run("rounds milliseconds with SecondsDecimalDigits", func(t *testing.T) {
		entries := []struct {
			input float64
			want  string
		}{
			{180000, "3 minutes"},
			{179999, "2 minutes 59 seconds"},
			{31536000000, "1 year"},
			{31535999999, "364 days 23 hours 59 minutes 59 seconds"},
			{86400000, "1 day"},
			{86399999, "23 hours 59 minutes 59 seconds"},
			{3600000, "1 hour"},
			{3599999, "59 minutes 59 seconds"},
			{7200000, "2 hours"},
			{7199999, "1 hour 59 minutes 59 seconds"},
		}

		for _, entry := range entries {
			got := Humanize(entry.input, Options{
				Verbose:                  true,
				SecondsDecimalDigits:     0,
				WithSecondsDecimalDigits: true,
			})

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})

	t.Run("have ColonNotation option", func(t *testing.T) {
		t.Run("With default option", func(t *testing.T) {
			entries := []struct {
				input float64
				want  string
			}{
				{1000, "0:01"},
				{1543, "0:01.5"},
				{60000, "1:00"},
				{90000, "1:30"},
				{95543, "1:35.5"},
				{600543, "10:00.5"},
				{3599543, "59:59.5"},
				{57599543, "15:59:59.5"},
			}

			for _, entry := range entries {
				got := Humanize(entry.input, Options{
					ColonNotation: true,
				})

				if got != entry.want {
					t.Errorf("Expected %v to be equal %v", got, entry.want)
				}
			}
		})

		t.Run("With SecondsDecimalDigits option", func(t *testing.T) {
			entries := []struct {
				input                float64
				secondsDecimalDigits int
				want                 string
			}{
				{1543, 0, "0:01"},
				{1543, 1, "0:01.5"},
				{1543, 2, "0:01.54"},
				{1543, 3, "0:01.543"},
				{95543, 0, "1:35"},
				{95543, 1, "1:35.5"},
				{95543, 2, "1:35.54"},
				{95543, 3, "1:35.543"},
				{600543, 3, "10:00.543"},
				{57599543, 3, "15:59:59.543"},
			}

			for _, entry := range entries {
				got := Humanize(entry.input, Options{
					ColonNotation:            true,
					SecondsDecimalDigits:     entry.secondsDecimalDigits,
					WithSecondsDecimalDigits: true,
				})

				if got != entry.want {
					t.Errorf("Expected %v to be equal %v", got, entry.want)
				}
			}
		})

		t.Run("With UnitCount and SecondsDecimalDigits option", func(t *testing.T) {
			entries := []struct {
				input                float64
				unitCount            int
				secondsDecimalDigits int
				want                 string
			}{
				{90000, 1, 0, "1"},
				{90000, 2, 0, "1:30"},
				{5400000, 3, 0, "1:30:00"},
				{95543, 2, 1, "1:35.5"},
				{3695543, 3, 1, "1:01:35.5"},
			}

			for _, entry := range entries {
				got := Humanize(entry.input, Options{
					ColonNotation:            true,
					UnitCount:                entry.unitCount,
					WithUnitCount:            true,
					SecondsDecimalDigits:     entry.secondsDecimalDigits,
					WithSecondsDecimalDigits: true,
				})

				if got != entry.want {
					t.Errorf("Expected %v to be equal %v", got, entry.want)
				}
			}
		})

		t.Run("With incompatible options", func(t *testing.T) {
			got := Humanize(3599543, Options{
				ColonNotation:         true,
				FormatSubMilliseconds: true,
			})
			want := "59:59.5"
			if got != want {
				t.Errorf("Expected %v to be equal %v", got, want)
			}

			got = Humanize(3599543, Options{
				ColonNotation:        true,
				SeparateMilliseconds: true,
			})

			if got != want {
				t.Errorf("Expected %v to be equal %v", got, want)
			}

			got = Humanize(3599543, Options{
				ColonNotation: true,
				Verbose:       true,
			})

			if got != want {
				t.Errorf("Expected %v to be equal %v", got, want)
			}

			got = Humanize(3599543, Options{
				ColonNotation: true,
				Compact:       true,
			})

			if got != want {
				t.Errorf("Expected %v to be equal %v", got, want)
			}
		})
	})
}
