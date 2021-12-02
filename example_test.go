package sauce_test

import (
	"embed"
	"log"
)

var (
	//go:embed static/*
	static embed.FS
)

func Example() {
	// print about the file
	file, err := static.Open("static/sauce.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
