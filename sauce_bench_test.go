package sauce_test

import (
	"bytes"
	"fmt"
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
	for b.Loop() {
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
	for range 10 {
		largeFile = append(largeFile, raw...)
	}

	b.ResetTimer()
	for b.Loop() {
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
	for b.Loop() {
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
	for b.Loop() {
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
	for b.Loop() {
		reader := bytes.NewReader(raw)
		_, _ = sauce.Read(reader)
	}
}

// Benchmark component functions

// BenchmarkUnsignedBinary1 measures single-byte binary parsing.
func BenchmarkUnsignedBinary1(b *testing.B) {
	testData := [1]byte{0xFF}

	b.ResetTimer()
	for b.Loop() {
		_ = layout.UnsignedBinary1(testData)
	}
}

// BenchmarkUnsignedBinary2 measures two-byte binary parsing.
func BenchmarkUnsignedBinary2(b *testing.B) {
	testData := [2]byte{0xFF, 0x00}

	b.ResetTimer()
	for b.Loop() {
		_ = layout.UnsignedBinary2(testData)
	}
}

// BenchmarkUnsignedBinary4 measures four-byte binary parsing.
func BenchmarkUnsignedBinary4(b *testing.B) {
	testData := [4]byte{0xFF, 0x00, 0x00, 0x00}

	b.ResetTimer()
	for b.Loop() {
		_ = layout.UnsignedBinary4(testData)
	}
}

// BenchmarkCommentByLine measures comment parsing by fixed 64-byte lines.
func BenchmarkCommentByLine(b *testing.B) {
	// Create test comment data: 3 lines of 64 bytes each
	commentData := []byte(strings.Repeat("A", 64) + strings.Repeat("B", 64) + strings.Repeat("C", 48))

	b.ResetTimer()
	for b.Loop() {
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
	for b.Loop() {
		_ = layout.CommentByLine(commentData)
	}
}

// BenchmarkCommentByBreak measures comment parsing by line breaks.
func BenchmarkCommentByBreak(b *testing.B) {
	// Create test comment data with line breaks
	commentData := []byte("Line 1\nLine 2\nLine 3\nLine 4\nLine 5")

	b.ResetTimer()
	for b.Loop() {
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
	for b.Loop() {
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
	for b.Loop() {
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
	for b.Loop() {
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
	for b.Loop() {
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
	for b.Loop() {
		_ = sauce.Contains(raw)
	}
}

// BenchmarkFileSizePerformance measures parsing performance with different file sizes.
func BenchmarkFileSizePerformance(b *testing.B) {
	baseFile, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		b.Fatalf("failed to load test file: %v", err)
	}

	sizes := []int{1, 10, 100, 1000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size_%dkB", size), func(b *testing.B) {
			// Create a file of the target size by repeating the base file
			largeFile := make([]byte, 0, size*1024)
			for len(largeFile) < size*1024 {
				largeFile = append(largeFile, baseFile...)
			}
			// Add SAUCE record at the end
			largeFile = append(largeFile, []byte("SAUCE00")...)
			// Add some dummy SAUCE data
			largeFile = append(largeFile, bytes.Repeat([]byte{0}, 128)...)

			b.ResetTimer()
			for b.Loop() {
				_ = sauce.Decode(largeFile)
			}
		})
	}
}

// BenchmarkMemoryAllocation measures memory allocation patterns.
func BenchmarkMemoryAllocation(b *testing.B) {
	raw, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		b.Fatalf("failed to load test file: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		_ = sauce.Decode(raw)
	}
}

// BenchmarkJSONSerializationSize measures JSON serialization with different record sizes.
func BenchmarkJSONSerializationSize(b *testing.B) {
	baseFile, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		b.Fatalf("failed to load test file: %v", err)
	}

	// Test with different comment sizes
	commentSizes := []int{0, 1, 10, 100}
	for _, size := range commentSizes {
		b.Run(fmt.Sprintf("Comments_%d", size), func(b *testing.B) {
			// Create a record with varying comment sizes
			rec := sauce.Decode(baseFile)

			// Create a new record with modified comment count
			originalRec := rec
			rec.Comnt.Count = size % 256 // Modify count for this test

			b.ResetTimer()
			for b.Loop() {
				_, _ = rec.JSON()
			}

			// Restore original record
			rec = originalRec
		})
	}
}

// BenchmarkConcurrentParsing measures performance under concurrent access.
func BenchmarkConcurrentParsing(b *testing.B) {
	raw, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		b.Fatalf("failed to load test file: %v", err)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = sauce.Decode(raw)
		}
	})
}

// BenchmarkInvalidRecords measures performance with invalid/malformed records.
func BenchmarkInvalidRecords(b *testing.B) {
	// Test with empty data
	b.Run("Empty", func(b *testing.B) {
		emptyData := []byte{}
		b.ResetTimer()
		for b.Loop() {
			_ = sauce.Decode(emptyData)
		}
	})

	// Test with incomplete SAUCE record
	b.Run("Incomplete", func(b *testing.B) {
		incompleteData := []byte("This is some dataSAUCE00")
		b.ResetTimer()
		for b.Loop() {
			_ = sauce.Decode(incompleteData)
		}
	})

	// Test with corrupted SAUCE record
	b.Run("Corrupted", func(b *testing.B) {
		corruptedData := []byte("This is some dataCORRUPT00")
		b.ResetTimer()
		for b.Loop() {
			_ = sauce.Decode(corruptedData)
		}
	})
}
