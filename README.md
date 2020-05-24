# pretty-milliseconds
Convert milliseconds to a human readable string: `1337000000` → `15d 11h 23m 20s`

![build](https://github.com/fernandoporazzi/pretty-milliseconds/workflows/build/badge.svg)

## Install 

```
$ go get -u github.com/fernandoporazzi/pretty-milliseconds
```

## API

func Humanize
```go
func Humanize(m float64, options ...Options) string
```

Humanize takes a float64 as milliseconds and, based on the options provided, returns a human readable string.

`Ex.: 144000000 -> 1d 16h 0m 0s`

```go
type Options struct {
	// Number of digits to appear after the seconds decimal point.
	SecondsDecimalDigits int

	// Whenever SecondsDecimalDigits is used, this has to be set to true
	WithSecondsDecimalDigits bool

	// Number of digits to appear after the milliseconds decimal point.
	MillisecondsDecimalDigits int

	// Keep milliseconds on whole seconds: 13s → 13.0s.
	// Useful when you are showing a number of seconds spent on an operation and don't want the width of the output to change when hitting a whole number.
	KeepDecimalsOnWholeSeconds bool

	// Only show the first unit: 1h 10m → 1h.
	// Also ensures that millisecondsDecimalDigits and secondsDecimalDigits are both set to 0.
	Compact bool

	// Number of units to show. Setting compact to true overrides this option.
	UnitCount int

	// WithUnitCount uses toggles wheter or not to use the UnitCount
	WithUnitCount bool

	// Use full-length units: 5h 1m 45s → 5 hours 1 minute 45 seconds
	Verbose bool

	// Show milliseconds separately. This means they won't be included in the decimal part of the seconds.
	SeparateMilliseconds bool

	// Show microseconds and nanoseconds.
	FormatSubMilliseconds bool

	// Display time using colon notation: 5h 1m 45s → 5:01:45. Always shows time in at least minutes: 1s → 0:01
	// Useful when you want to display time without the time units, similar to a digital watch.
	// Setting colonNotation to true overrides the following options to false:
	// Compact
	// FormatSubMilliseconds
	// SeparateMilliseconds
	// Verbose
	ColonNotation bool
}
```

## Usage

```go
package main

import (
	"fmt"

	prettyms "github.com/fernandoporazzi/pretty-milliseconds"
)

func main() {
	r := prettyms.Humanize(980000000, prettyms.Options{
		Verbose: true,
	})

	fmt.Println(r) // 11 days 8 hours 13 minutes 20 seconds

	r = prettyms.Humanize(1337000000, prettyms.Options{
		ColonNotation: true,
	})

	fmt.Println(r) // => 15:11:23:20

	r = prettyms.Humanize(72360000000, prettyms.Options{
		Compact: true,
	})

	fmt.Println(r) // => 2y

	r = prettyms.Humanize(72360000000, prettyms.Options{
		Compact: true,
		Verbose: true,
	})

	fmt.Println(r) // => 2 years

	r = prettyms.Humanize(0.123456789, prettyms.Options{
		FormatSubMilliseconds: true,
	})

	fmt.Println(r) // => 123µs 456ns

	r = prettyms.Humanize(33333, prettyms.Options{
		SecondsDecimalDigits:     2,
		WithSecondsDecimalDigits: true,
	})

	fmt.Println(r) // => 33.33s

	r = prettyms.Humanize(2000, prettyms.Options{
		SecondsDecimalDigits:     0,
		WithSecondsDecimalDigits: true,
		Verbose:                  true,
	})

	fmt.Println(r) // => 2 seconds

	r = prettyms.Humanize(60000, prettyms.Options{
		UnitCount:     0,
		WithUnitCount: true,
	})

	fmt.Println(r) // => 1m

	r = prettyms.Humanize(4020000, prettyms.Options{
		UnitCount:     1,
		WithUnitCount: true,
	})

	fmt.Println(r) // => 1h

	r = prettyms.Humanize(4020000, prettyms.Options{
		UnitCount:     2,
		WithUnitCount: true,
	})

	fmt.Println(r) // => 1h 7m

	r = prettyms.Humanize(44863200000, prettyms.Options{
		UnitCount:     3,
		WithUnitCount: true,
	})

	fmt.Println(r) // => 1y 154d 6h
}
```


## License

[MIT License](https://github.com/fernandoporazzi/pretty-milliseconds/blob/master/LICENSE)
