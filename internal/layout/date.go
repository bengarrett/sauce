// The date the file was created. The format for the date is CCYYMMDD (century, year, month, day).
// Example: 4 May 2013 would be stored as "20130504".
// See http://www.acid.org/info/sauce/sauce.htm
package layout

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

var (
	ErrParseDate = errors.New("parse date to integer conversion")
	ErrSauceDate = errors.New("sauce date error")
)

// Dates in multiple output formats.
type Dates struct {
	Value string    `json:"value" xml:"value"`
	Time  time.Time `json:"iso" xml:"date"`
	Epoch int64     `json:"epoch" xml:"epoch,attr"`
}

func (d *Layout) Dates() Dates {
	t, err := d.parseDate()
	if err != nil {
		log.Printf("%s: %s\n", ErrSauceDate, err)
		return Dates{}
	}
	u := t.Unix()
	return Dates{
		Value: d.Date.String(),
		Time:  t,
		Epoch: u,
	}
}

func (d *Layout) parseDate() (t time.Time, err error) {
	da := d.Date
	dy, err := strconv.Atoi(string(da[0:4]))
	if err != nil {
		return t, fmt.Errorf("year failed: %v: %w", dy, ErrParseDate)
	}
	dm, err := strconv.Atoi(string(da[4:6]))
	if err != nil {
		return t, fmt.Errorf("month failed: %v: %w", dm, ErrParseDate)
	}
	dd, err := strconv.Atoi(string(da[6:8]))
	if err != nil {
		return t, fmt.Errorf("day failed: %v: %w", dd, ErrParseDate)
	}
	return time.Date(dy, time.Month(dm), dd, 0, 0, 0, 0, time.UTC), nil
}
