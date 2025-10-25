package layout

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

// The date the file was created. The format for the date is CCYYMMDD (century, year, month, day).
// Example: 4 May 2013 would be stored as "20130504".
// See http://www.acid.org/info/sauce/sauce.htm

var (
	ErrParseDate = errors.New("parse date to integer conversion")
	ErrSauceDate = errors.New("sauce date error")
)

// Dates is the date the file was created, in multiple time formats.
type Dates struct {
	Value string    `json:"value" xml:"value"`      // date format using CCYYMMDD (century, year, month, day)
	Time  time.Time `json:"iso" xml:"date"`         // time as a go time type
	Epoch int64     `json:"epoch" xml:"epoch,attr"` // epoch unix time, is the number of seconds since 1 Jan 1970
}

func (d *Layout) Dates() Dates {
	tt, err := d.date()
	if err != nil {
		log.Printf("%s: %s\n", ErrSauceDate, err)
		return Dates{}
	}
	epoch := tt.Unix()
	return Dates{
		Value: d.Date.String(),
		Time:  tt,
		Epoch: epoch,
	}
}

func (d *Layout) date() (time.Time, error) {
	dd := d.Date
	year, err := strconv.Atoi(string(dd[0:4]))
	if err != nil {
		return time.Time{},
			fmt.Errorf("year failed: %v: %w", year, ErrParseDate)
	}
	month, err := strconv.Atoi(string(dd[4:6]))
	if err != nil {
		return time.Time{},
			fmt.Errorf("month failed: %v: %w", month, ErrParseDate)
	}
	day, err := strconv.Atoi(string(dd[6:8]))
	if err != nil {
		return time.Time{},
			fmt.Errorf("day failed: %v: %w", day, ErrParseDate)
	}
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}
