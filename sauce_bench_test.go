package sauce_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/bengarrett/sauce"
	"github.com/bengarrett/sauce/internal/layout"
)

// Benchmark hot path functions: Index, Decode, Trim

// BenchmarkIndex measures the backwards search for SAUCE record position.
func BenchmarkIndex(b *testing.B) {
	raw, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		b.Fatalf("failed to load test file: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sauce.Index(raw)
	}
}

// BenchmarkIndexLargeFile tests index search on a larger file to stress backwards search.
func BenchmarkIndexLargeFile(b *testing.B) {
	raw, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		b.Fatalf("failed to load test file: %v", err)
	}

	// Create a larger file by repeating the content
	largeFile := make([]byte, 0, len(raw)*10)
	for j := 0; j < 10; j++ {
		largeFile = append(largeFile, raw...)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sauce.Index(largeFile)
	}
}

// BenchmarkDecode measures full SAUCE record parsing (main API).
func BenchmarkDecode(b *testing.B) {
	raw, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		b.Fatalf("failed to load test file: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sauce.Decode(raw)
	}
}

// BenchmarkTrim measures SAUCE metadata removal.
func BenchmarkTrim(b *testing.B) {
	raw, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		b.Fatalf("failed to load test file: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sauce.Trim(raw)
	}
}

// BenchmarkRead measures complete file read and parse operation.
func BenchmarkRead(b *testing.B) {
	raw, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		b.Fatalf("failed to load test file: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := bytes.NewReader(raw)
		_, _ = sauce.Read(reader)
	}
}

// Benchmark component functions

// BenchmarkUnsignedBinary1 measures single-byte binary parsing.
func BenchmarkUnsignedBinary1(b *testing.B) {
	testData := [1]byte{0xFF}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = layout.UnsignedBinary1(testData)
	}
}

// BenchmarkUnsignedBinary2 measures two-byte binary parsing.
func BenchmarkUnsignedBinary2(b *testing.B) {
	testData := [2]byte{0xFF, 0x00}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = layout.UnsignedBinary2(testData)
	}
}

// BenchmarkUnsignedBinary4 measures four-byte binary parsing.
func BenchmarkUnsignedBinary4(b *testing.B) {
	testData := [4]byte{0xFF, 0x00, 0x00, 0x00}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = layout.UnsignedBinary4(testData)
	}
}

// BenchmarkCommentByLine measures comment parsing by fixed 64-byte lines.
func BenchmarkCommentByLine(b *testing.B) {
	// Create test comment data: 3 lines of 64 bytes each
	commentData := []byte(strings.Repeat("A", 64) + strings.Repeat("B", 64) + strings.Repeat("C", 48))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = layout.CommentByLine(commentData)
	}
}

// BenchmarkCommentByLineLarge measures comment parsing with many lines.
func BenchmarkCommentByLineLarge(b *testing.B) {
	// Create test comment data: 255 lines of 64 bytes each (max SAUCE comments)
	commentData := make([]byte, 255*64)
	for i := range commentData {
		commentData[i] = byte(i % 256)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = layout.CommentByLine(commentData)
	}
}

// BenchmarkCommentByBreak measures comment parsing by line breaks.
func BenchmarkCommentByBreak(b *testing.B) {
	// Create test comment data with line breaks
	commentData := []byte("Line 1\nLine 2\nLine 3\nLine 4\nLine 5")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = layout.CommentByBreak(commentData)
	}
}

// Benchmark JSON/XML serialization

// BenchmarkJSON measures JSON serialization.
func BenchmarkJSON(b *testing.B) {
	raw, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		b.Fatalf("failed to load test file: %v", err)
	}
	rec := sauce.Decode(raw)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rec.JSON()
	}
}

// BenchmarkJSONIndent measures indented JSON serialization.
func BenchmarkJSONIndent(b *testing.B) {
	raw, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		b.Fatalf("failed to load test file: %v", err)
	}
	rec := sauce.Decode(raw)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rec.JSONIndent("    ")
	}
}

// BenchmarkXML measures XML serialization.
func BenchmarkXML(b *testing.B) {
	raw, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		b.Fatalf("failed to load test file: %v", err)
	}
	rec := sauce.Decode(raw)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rec.XML()
	}
}

// BenchmarkXMLIndent measures indented XML serialization.
func BenchmarkXMLIndent(b *testing.B) {
	raw, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		b.Fatalf("failed to load test file: %v", err)
	}
	rec := sauce.Decode(raw)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rec.XMLIndent("  ")
	}
}

// Benchmark Contains helper function

// BenchmarkContains measures the helper function that checks for SAUCE presence.
func BenchmarkContains(b *testing.B) {
	raw, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		b.Fatalf("failed to load test file: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sauce.Contains(raw)
	}
}
