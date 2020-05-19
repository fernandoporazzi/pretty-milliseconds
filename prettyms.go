package prettyms

import (
	parsems "github.com/fernandoporazzi/parse-ms"
)

// Humanize takes a float64 as milliseconds and
// returns a human readable string
//
// Ex.: 144000000 -> 1d 16h 0m 0s
func Humanize(m float64) parsems.ParsedMilliseconds {
	return parsems.Parse(float64(1337000000))
}
