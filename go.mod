module github.com/bengarrett/sauce

// When updating go version, also update .github/workflows/go.yml
go 1.26.3

require golang.org/x/text v0.37.0

require (
	github.com/klauspost/compress v1.18.0 // indirect
	go.uber.org/nilaway v0.0.0-20251021214447-34f56b8c16b9 // indirect
	golang.org/x/mod v0.35.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/tools v0.44.0 // indirect
)

tool go.uber.org/nilaway/cmd/nilaway
