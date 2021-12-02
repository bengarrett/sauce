package record_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/bengarrett/sauce/internal/record"
)

func Test_data_dates(t *testing.T) {
	var (
		empty   record.Date
		empties record.Dates
	)
	tests := []struct {
		name string
		date record.Date
		want record.Dates
	}{
		{"empty", empty, empties},
		{"example", exampleData().Date, record.Dates{
			Value: "20161126",
			Time:  time.Date(2016, 11, 26, 0, 0, 0, 0, time.UTC),
			Epoch: 1480118400,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &record.Data{
				Date: tt.date,
			}
			if got := d.Dates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("data.dates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_data_DataType(t *testing.T) {
	tests := []struct {
		name     string
		datatype record.DataType
		want     record.Datas
	}{
		{"none", [1]byte{0},
			record.Datas{
				Type: record.Nones,
				Name: record.Nones.String()}},
		{"archive", [1]byte{7},
			record.Datas{
				Type: record.Archives,
				Name: record.Archives.String()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &record.Data{
				Datatype: tt.datatype,
			}
			if got := d.DataType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("data.DataType() = %v, want %v", got, tt.want)
			}
		})
	}
}
