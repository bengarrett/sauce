package layout_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/bengarrett/sauce/internal/layout"
)

func Test_data_dates(t *testing.T) {
	var (
		empty   layout.Date
		empties layout.Dates
	)
	tests := []struct {
		name string
		date layout.Date
		want layout.Dates
	}{
		{"empty", empty, empties},
		{"example", exampleData().Date, layout.Dates{
			Value: "20161126",
			Time:  time.Date(2016, 11, 26, 0, 0, 0, 0, time.UTC),
			Epoch: 1480118400,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &layout.Layout{
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
		datatype layout.DataType
		want     layout.Datas
	}{
		{"none", [1]byte{0},
			layout.Datas{
				Type: layout.Nones,
				Name: layout.Nones.String()}},
		{"archive", [1]byte{7},
			layout.Datas{
				Type: layout.Archives,
				Name: layout.Archives.String()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &layout.Layout{
				Datatype: tt.datatype,
			}
			if got := d.DataType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("data.DataType() = %v, want %v", got, tt.want)
			}
		})
	}
}
