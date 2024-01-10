// The code in humanize/bytes.go is derived from [labstack/gommon], Common packages for Go
//
// [The MIT License] (MIT) Copyright (c) 2018 labstack
//
// [labstack/gommon]: https://github.com/labstack/gommon
// [The MIT License]: https://github.com/labstack/gommon/blob/master/LICENSE

// Package humanize is parses some limited time and byte sizes data to human readable formats.
package humanize

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const (
	_ = 1.0 << (binaryBase * iota)
	kib
	mib
	gib
	tib
	pib

	oneDecimalPoint  = "%.1f %s"
	twoDecimalPoints = "%.2f %s"
	binaryBase       = 10
	kb               = 1000
	mb               = kb * kb
	gb               = mb * kb
	tb               = gb * kb
	pb               = tb * kb
)

// Binary formats the bytes integer to localized, readable string using binary units of measure.
func Binary(b int64, t language.Tag) string {
	p := message.NewPrinter(t)
	value := float64(b)
	var multiple string
	switch {
	case b >= pib:
		value /= pib
		multiple = "PiB"
	case b >= tib:
		value /= tib
		multiple = "TiB"
	case b >= gib:
		value /= gib
		multiple = "GiB"
	case b >= mib:
		value /= mib
		multiple = "MiB"
	case b >= kib:
		value /= kib
		return p.Sprintf(oneDecimalPoint, value, "KiB")
	case b == 0:
		return "0"
	default:
		return p.Sprintf("%dB", b)
	}
	return p.Sprintf(twoDecimalPoints, value, multiple)
}

// Decimal formats the bytes integer to localized readable string using decimal units of measure.
func Decimal(b int64, t language.Tag) string {
	p := message.NewPrinter(t)
	value := float64(b)
	var multiple string
	switch {
	case b >= pb:
		value /= pb
		multiple = "PB"
	case b >= tb:
		value /= tb
		multiple = "TB"
	case b >= gb:
		value /= gb
		multiple = "GB"
	case b >= mb:
		value /= mb
		multiple = "MB"
	case b >= kb:
		value /= kb
		return p.Sprintf(oneDecimalPoint, value, "kB")
	case b == 0:
		return "0"
	default:
		return p.Sprintf("%dB", b)
	}
	return p.Sprintf(twoDecimalPoints, value, multiple)
}
