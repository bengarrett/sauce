# Copilot Instructions for Sauce Repository

## Project Overview

**Sauce** is a Go module that parses SAUCE (Standard Architecture for Universal Comment Extensions) metadata embedded in files, primarily used for ANSI art. The project reads and decodes SAUCE records from file headers to extract metadata like title, author, group, date, file type information, and comments.

## Build, Test & Lint Commands

### Using Task (Recommended)
Task is the primary build tool for this repository. Run `task --list` to see all available tasks.

**Common Commands:**
- `task test` - Run the test suite
- `task testr` - Run tests with race detection (slower but more thorough)
- `task lint` - Format code with `gofumpt` and lint with `golangci-lint`
- `task nil` - Run nilaway static analysis for nil dereferences
- `task pkg-update` - Update all dependencies to latest versions
- `task pkg-patch` - Apply patch-level dependency updates

**Using Go Directly:**
- `go test -count 1 ./...` - Run all tests
- `go test -count 1 -race ./...` - Run tests with race detection
- `gofumpt -l -w .` - Format all Go files

### Running Single Tests
```bash
go test -count 1 -run TestTrim ./...  # Run specific test by function name
go test -count 1 -v ./...              # Run with verbose output to see individual tests
```

### Linting Configuration
- Linter: `golangci-lint` configured in `.golangci.yml`
- Formatter: `gofumpt` with `goimports`, `gci`, `gofmt`
- Static analysis: `nilaway` (enabled as a go tool)
- Max cyclomatic complexity: 15
- Note: Several linters are disabled (varnamelen, wsl, depguard, etc.) - these are intentionally disabled

## Architecture & Code Organization

### High-Level Design
1. **Public API** (`sauce.go`) - Core functions:
   - `Read(io.Reader)` - Reads SAUCE from a file/reader
   - `Decode([]byte)` - Decodes raw bytes into a Record
   - `Contains([]byte)` - Checks if SAUCE data exists in bytes
   - `Index([]byte)` - Finds SAUCE position in data
   - `Trim([]byte)` - Removes SAUCE metadata from data

2. **Internal Layout** (`internal/layout/`) - Low-level binary structure handling:
   - Defines the exact byte layout of SAUCE records according to specification
   - Contains type-specific parsers for different data types (character art, bitmap, audio, executable, etc.)
   - Handles binary parsing of fixed-size fields and flags
   - Each data type (Character, Bitmap, Audio, Archive, Executable, etc.) has its own subpackage with type-specific info parsing

3. **Humanize** (`humanize/`) - User-friendly data formatting:
   - Converts raw SAUCE values into human-readable strings
   - Handles date formatting and byte size formatting (decimal and binary)

### Data Flow
- **Input**: File bytes → 
- **Layout Layer**: Extracts fixed-position binary fields (128-byte SAUCE record, optional COMNT comment block) →
- **Record Layer**: Creates structured `Record` with decoded metadata →
- **Output**: Public API methods for JSON/XML export or direct field access

### Key Struct: `Record`
Located in `sauce.go`, contains:
- `ID`, `Version` - Identifiers
- `Title`, `Author`, `Group` - Metadata strings
- `Date` - Parsed date with Unix timestamp
- `FileSize` - Bytes in decimal and binary notation
- `DataType`, `FileType` - Type classification with descriptions
- `TypeInfo` - Type-specific numerical info (1-4) and string info
- `Flags` - Type-specific flag interpretations
- `FontName` - Font preference
- `Comnt` - Optional comment block structure

## Key Conventions

### Constants
- `SAUCE` - Fixed 5-byte ID that marks valid records
- `"00"` - Version is always "00" per SAUCE spec
- `Date = "20060102"` - SAUCE date format constant (CCYYMMDD) using Go time layout
- `EOF = 26` (SUB character) - Optional end-of-file marker before SAUCE record
- Comment line size is fixed at 64 bytes (`ComntLineSize`)

### Package Structure
- `layout` package types are unexported (lowercase) - implementation detail
- The public `Record` struct wraps layout types with user-friendly accessors
- Tests use embedded `static` file system to load test fixtures (see `sauce_test.go`)

### Layout Constants (in `internal/layout`)
- SAUCE record is always 128 bytes
- Comment block header is 10 bytes per line of comments
- Each comment line is exactly 64 bytes
- Maximum 255 comment lines per block

### Testing Patterns
- Test files use `static.ReadFile()` to load fixtures from `static/` directory (embedded with `go:embed`)
- Tests check both positive cases (valid SAUCE) and negative cases (missing/malformed SAUCE)
- Use `reflect.DeepEqual()` for struct comparison in tests
- Tests verify field values, JSON/XML marshaling, and trim operations

### Code Style
- Use `strings.TrimSpace()` for field cleanup (removes padding nulls)
- Type-specific info uses array indices (1-4) matching SAUCE specification
- Error handling distinguishes between missing SAUCE and malformed data
- JSON/XML marshaling uses appropriate field tags for API consistency

### Type-Specific Handling
Different data types have different interpretations of the 4 TInfo fields and flags. The `layout` subpackage determines type-specific meanings (e.g., TInfo1 might mean "width" for characters, "pixel width" for graphics, "sample rate" for audio, etc.). Always reference the SAUCE specification at http://www.acid.org/info/sauce/sauce.htm for type definitions.
