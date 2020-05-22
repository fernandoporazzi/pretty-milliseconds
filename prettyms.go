package prettyms

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	parsems "github.com/fernandoporazzi/parse-ms"
)

// Result holds an array with values, such as year, days, hours and so on...
type result struct {
	Values []string
}

var options Options

func newResult() *result {
	return &result{}
}

func pluralize(l string, v float64) string {
	if v == 1 {
		return l
	}

	return l + "s"
}

func floorDecimals(value float64, decimalDigits int) string {
	secondRoundingEpsilon := 0.0000001
	pow10 := math.Pow10(decimalDigits)
	v := float64(value)
	flooredInterimValue := math.Floor((v * pow10) + secondRoundingEpsilon)

	flooredValue := math.Round(flooredInterimValue) / pow10

	return fmt.Sprintf("%.[2]*[1]f", flooredValue, decimalDigits)
}

// Append pushes a new slice into the r.Values array
func (r *result) Append(value float64, args ...string) {
	long := args[0]
	short := args[1]

	var valueString string

	if len(args) == 3 {
		valueString = args[2]
	}

	if (len(r.Values) == 0 || !options.ColonNotation) && value == 0 && !(options.ColonNotation && short == "m") {
		return
	}

	if value > 0 {
		valueString = strconv.FormatFloat(value, 'f', -1, 64)

		if options.Compact && options.Verbose {
			valueString = fmt.Sprintf("%d", int(value))
		}
	} else {
		valueString = "0"
	}

	var prefix, suffix string

	if options.ColonNotation {
		if len(r.Values) > 0 {
			prefix = ":"
		} else {
			prefix = ""
		}
		suffix = ""

		var wholeDigits, minLength int

		if strings.Contains(valueString, ".") {
			s := strings.Split(valueString, ".")[0]
			wholeDigits = len(s)
		} else {
			wholeDigits = len(valueString)
		}

		if len(r.Values) > 0 {
			minLength = 2
		} else {
			minLength = 1
		}

		repeat := int(math.Max(0, float64(minLength-wholeDigits)))
		valueString = strings.Repeat("0", repeat) + valueString
	} else {
		prefix = ""
		if options.Verbose {
			suffix = " " + pluralize(long, value)
		} else {
			suffix = short
		}
	}

	r.Values = append(r.Values, prefix+valueString+suffix)
}

// Humanize takes a float64 as milliseconds and
// returns a human readable string
//
// Ex.: 144000000 -> 1d 16h 0m 0s
func Humanize(m float64, o ...Options) string {
	r := newResult()

	if len(o) == 1 {
		options = o[0]
	}

	if options.ColonNotation {
		options.Compact = false
		options.FormatSubMilliseconds = false
		options.SeparateMilliseconds = false
		options.Verbose = false
	}

	if options.Compact {
		options.SecondsDecimalDigits = 0
		options.MillisecondsDecimalDigits = 0
	}

	parsed := parsems.Parse(m)

	r.Append(math.Trunc(parsed.Days/365), "year", "y")
	r.Append(float64(int(parsed.Days)%365), "day", "d")
	r.Append(parsed.Hours, "hour", "h")
	r.Append(parsed.Minutes, "minute", "m")

	if options.SeparateMilliseconds || options.FormatSubMilliseconds || m < 1000 {
		r.Append(parsed.Seconds, "second", "s")

		if options.FormatSubMilliseconds {
			r.Append(parsed.Milliseconds, "millisecond", "ms")
			r.Append(parsed.Microseconds, "microsecond", "µs")
			r.Append(parsed.Nanoseconds, "nanosecond", "ns")
		} else {
			millisecondsAndBelow := parsed.Milliseconds + (parsed.Microseconds / 1000) + (parsed.Nanoseconds / 1000000)

			millisecondsDecimalDigits := options.MillisecondsDecimalDigits

			var roundedMiliseconds float64
			if millisecondsAndBelow >= 1 {
				roundedMiliseconds = math.Round(millisecondsAndBelow)
			} else {
				roundedMiliseconds = math.Ceil(millisecondsAndBelow)
			}

			var millisecondsString string
			if millisecondsDecimalDigits != 0 {
				millisecondsString = fmt.Sprintf("%.[2]*[1]f", millisecondsAndBelow, millisecondsDecimalDigits)
			} else {
				millisecondsString = fmt.Sprintf("%f", roundedMiliseconds)
			}

			pf, err := strconv.ParseFloat(millisecondsString, 64)
			if err != nil {
				panic(err)
			}
			r.Append(pf, "millisecond", "ms", millisecondsString)
		}
	} else {
		seconds := math.Mod(m/1000, 60)

		secondsDecimalDigits := 1
		if options.WithSecondsDecimalDigits {
			secondsDecimalDigits = options.SecondsDecimalDigits
		}

		secondsFixed := floorDecimals(seconds, secondsDecimalDigits)

		f, err := strconv.ParseFloat(secondsFixed, 64)
		if err != nil {
			panic(err)
		}

		secondsString := strconv.FormatFloat(f, 'f', -1, 64)
		if options.KeepDecimalsOnWholeSeconds {
			secondsString = secondsFixed
		}

		f, err = strconv.ParseFloat(secondsString, 64)
		if err != nil {
			panic(err)
		}

		r.Append(f, "second", "s", secondsString)
	}

	if len(r.Values) == 0 {
		if options.Verbose {
			return "0 milliseconds"
		}
		return "0ms"
	}

	if options.Compact {
		return r.Values[0]
	}

	if options.WithUnitCount {
		separator := " "
		if options.ColonNotation {
			separator = ""
		}

		max := int(math.Max(float64(options.UnitCount), 1))
		slice := r.Values[0:max]
		return strings.Join(slice[:], separator)
	}

	separator := " "
	if options.ColonNotation {
		separator = ""
	}

	return strings.Join(r.Values[:], separator)
}
