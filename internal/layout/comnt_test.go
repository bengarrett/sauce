package layout_test

import (
	"reflect"
	"testing"

	"github.com/bengarrett/sauce/internal/layout"
)

func sauceIndex() int {
	return layout.Index(raw())
}

func Test_record_Comnt(t *testing.T) {
	type args struct {
		count      layout.Comments
		sauceIndex int
	}
	tests := []struct {
		name       string
		r          layout.Data
		args       args
		wantLength int
	}{
		{
			"example", layout.Data(raw()),
			args{count: [1]byte{1}, sauceIndex: sauceIndex()},
			1 * layout.ComntLineSize,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBlock := tt.r.Comnt(tt.args.count, tt.args.sauceIndex); !reflect.DeepEqual(gotBlock.Length, tt.wantLength) {
				t.Errorf("layout.Comnt() = %v, want %v", gotBlock.Length, tt.wantLength)
			}
		})
	}
}

func Test_data_CommentBlock(t *testing.T) {
	tests := []struct {
		name string
		data layout.Layout
		want []string
	}{
		{"empty", layout.Layout{}, nil},
		{"example", exampleData(), []string{commentResult}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &layout.Comnt{
				Lines: tt.data.Comnt.Lines,
			}
			d := &layout.Layout{
				Comments: tt.data.Comments,
				Comnt:    *c,
			}
			if gotC := d.CommentBlock(); !reflect.DeepEqual(gotC.Comment, tt.want) {
				t.Errorf("data.CommentBlock() = %v, want %v", gotC.Comment, tt.want)
			}
		})
	}
}

func Test_CommentByBreak(t *testing.T) {
	tests := []struct {
		name      string
		b         []byte
		wantLines []string
	}{
		{"empty", []byte{}, nil},
		{"example", exampleData().Comnt.Lines, []string{commentResult}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLines := layout.CommentByBreak(tt.b); !reflect.DeepEqual(gotLines, tt.wantLines) {
				t.Errorf("CommentByBreak() = %v, want %v", gotLines, tt.wantLines)
			}
		})
	}
}

func Test_CommentByLine(t *testing.T) {
	tests := []struct {
		name      string
		b         []byte
		wantLines []string
	}{
		{"empty", []byte{}, nil},
		{"example", exampleData().Comnt.Lines, []string{commentResult}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLines := layout.CommentByLine(tt.b); !reflect.DeepEqual(gotLines, tt.wantLines) {
				t.Errorf("CommentByLine() = %v, want %v", gotLines, tt.wantLines)
			}
		})
	}
}
