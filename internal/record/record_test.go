package record_test

import (
	"embed"
	"log"

	"github.com/bengarrett/sauce/internal/record"
)

const (
	commentResult = "Any comments go here.                                           "
	example       = "static/sauce.txt"
)

var (
	//go:embed static/*
	static embed.FS
)

func raw() []byte {
	b, err := static.ReadFile(example)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func exampleData() record.Data {
	return record.Record(raw()).Extract()
}
