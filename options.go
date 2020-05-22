package prettyms

// Options struct allows consumer to overwrite default config values
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
