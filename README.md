# Package sauce

Package sauce is a [Go module](https://go.dev/) that parses SAUCE (Standard Architecture for
Universal Comment Extensions) metadata.

> The Standard Architecture for Universal Comment Extensions or SAUCE as it is
more commonly known, is an architecture or protocol for attaching meta data
or comments to files. Mainly intended for [ANSI art files](https://en.wikipedia.org/wiki/ANSI_art), SAUCE has always
had provisions for many different file types.

For the complete specification see http://www.acid.org/info/sauce/sauce.htm.

## Quick usage


[Go Package with docs and examples.](https://pkg.go.dev/github.com/bengarrett/sauce)

```go
// open a file
file, err := os.Open("artwork.ans")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

// read the file and create a SAUCE record
sr, err := sauce.NewRecord(file)
if err != nil {
    log.Println(err)
    return
}

// print specific SAUCE fields
fmt.Printf("%q\n", sr.Title)
fmt.Printf("Author:\t%s.\n", sr.Author)
fmt.Printf("Group:\t%s.\n", sr.Group)
fmt.Printf("Date:\t%s.\n", sr.Date.Time.Format(time.ANSIC))

// print the SAUCE data as indented JSON
js, err := sr.JSONIndent("  ")
if err != nil {
    log.Println(err)
    return
}
fmt.Printf("%s", js)
```