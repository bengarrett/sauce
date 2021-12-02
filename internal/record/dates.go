package record

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

var ErrDate = errors.New("parse date to integer conversion")

// Dates in multiple output formats.
type Dates struct {
	Value string    `json:"value" xml:"value"`
	Time  time.Time `json:"iso" xml:"date"`
	Epoch int64     `json:"epoch" xml:"epoch,attr"`
}

func (d *Data) Dates() Dates {
	t, err := d.parseDate()
	if err != nil {
		fmt.Printf("sauce date error: %s\n", err)
	}
	u := t.Unix()
	return Dates{
		Value: d.Date.String(),
		Time:  t,
		Epoch: u,
	}
}

func (d *Data) parseDate() (t time.Time, err error) {
	da := d.Date
	dy, err := strconv.Atoi(string(da[0:4]))
	if err != nil {
		return t, fmt.Errorf("year failed: %v: %w", dy, ErrDate)
	}
	dm, err := strconv.Atoi(string(da[4:6]))
	if err != nil {
		return t, fmt.Errorf("month failed: %v: %w", dm, ErrDate)
	}
	dd, err := strconv.Atoi(string(da[6:8]))
	if err != nil {
		return t, fmt.Errorf("day failed: %v: %w", dd, ErrDate)
	}
	return time.Date(dy, time.Month(dm), dd, 0, 0, 0, 0, time.UTC), nil
}
