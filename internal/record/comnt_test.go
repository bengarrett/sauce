package record_test

import (
	"reflect"
	"testing"

	"github.com/bengarrett/sauce/internal/record"
)

// func Test_record_comnt(t *testing.T) {
// 	type args struct {
// 		count      comments
// 		sauceIndex int
// 	}
// 	tests := []struct {
// 		name       string
// 		r          record
// 		args       args
// 		wantLength int
// 	}{
// 		{"example", record(raw()), args{count: [1]byte{1}, sauceIndex: sauceIndex()}, 1 * comntLineSize},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if gotBlock := tt.r.comnt(tt.args.count, tt.args.sauceIndex); !reflect.DeepEqual(gotBlock.length, tt.wantLength) {
// 				t.Errorf("record.comnt() = %v, want %v", gotBlock.length, tt.wantLength)
// 			}
// 		})
// 	}
// }

func Test_data_comment(t *testing.T) {
	tests := []struct {
		name string
		data record.Data
		want []string
	}{
		{"empty", record.Data{}, nil},
		{"example", exampleData(), []string{commentResult}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &record.Comnt{
				Lines: tt.data.Comnt.Lines,
			}
			d := &record.Data{
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
			if gotLines := record.CommentByBreak(tt.b); !reflect.DeepEqual(gotLines, tt.wantLines) {
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
			if gotLines := record.CommentByLine(tt.b); !reflect.DeepEqual(gotLines, tt.wantLines) {
				t.Errorf("CommentByLine() = %v, want %v", gotLines, tt.wantLines)
			}
		})
	}
}
