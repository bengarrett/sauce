[![Go Reference](https://pkg.go.dev/badge/github.com/bengarrett/sauce.svg)](https://pkg.go.dev/github.com/bengarrett/sauce) &nbsp; [![Go Report Card](https://goreportcard.com/badge/github.com/bengarrett/sauce)](https://goreportcard.com/report/github.com/bengarrett/sauce)

# Package sauce

Package sauce is a [Go module](https://go.dev/) that parses SAUCE (Standard Architecture for
Universal Comment Extensions) metadata.

> The Standard Architecture for Universal Comment Extensions or SAUCE as it is
more commonly known, is an architecture or protocol for attaching meta data
or comments to files. Mainly intended for [ANSI art files](https://en.wikipedia.org/wiki/ANSI_art), SAUCE has always
had provisions for many different file types.

For the complete specification see:<br>
http://www.acid.org/info/sauce/sauce.htm<br>
https://github.com/radman1/sauce<br>

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

## SAUCE as an API reference

- `id`<br>
SAUCE identification. This should be equal to `SAUCE`.

- `version`<br>
SAUCE version number, should be `00`.

- `title`<br>
Title of the file.

- `author`<br>
The nick, name or handle of the creator of the file.

- `group`<br>
The name of the group or company the creator is employed by.

- `date` - The date the file was created.
- - `value` - SAUCE date format, CCYYMMDD (century, year, month, day).
- - `iso` -  `value` as an ISO 8601 date.
- - `epoch`- `value` as Unix time, the number of seconds since 1 Jan 1970.

- `fileSize`
- - `bytes` - The reported file size not including the SAUCE information.
- - `decimal` - `bytes` returned as a base 10 value (kilo, mega...).
- - `binary` - `bytes` returned as a base 2 value (kibi, mebi...).

- `dataType` - Type of data.
- - `type` - `DataType` value.
- - `name` - `DataType` name.

- `fileType` - Type of file.
- - `type` - `FileType` value.
- - `name` - `FileType` name.

- `typeInfo` - Type dependant information, see http://www.acid.org/info/sauce/sauce.htm#FileType.
- - `1`
- - - `value` - Value of `TInfo1`.
- - - `info` - Human readable description of the value.
- - `2`
- - - `value` - Value of `TInfo2`.
- - - `info` - Human readable description of the value.
- - `3`
- - - `value` - Value of `TInfo3`.
- - - `info` - Human readable description of the value.

* - `flags` - Type dependant flags.
* - - `decimal` - Value as an unsigned integer.
* - - `binary` - Value in binary notation.
- - - `nonBlinkMode` - Request ANSi non-blink mode (iCE Color).
- - - - `flag` - Value as binary bit boolean ("0" or "1").
- - - - `interpretation` - Human readable description of the value.
* - - `letterSpacing` - ANSi letter-spacing to request 8 or 9 pixel font selection.
* - - - `flag` - Value as a 2-bit binary string ("00", "01", "10").
* - - - `interpretation` - Human readable description of the value.
- - - `aspectRatio` - ANSi aspect ratio to request LCD square or CRT monitor style pixels.
- - - - `flag` - Value as a 2-bit binary string ("00", "01", "10").
- - - - `interpretation` - Human readable description of the value.
* - `fontName` - The creator's preferred font to view the ANSi artwork, see http://www.acid.org/info/sauce/sauce.htm#FontName.

- `comments` - Comments or notes from the creator.
- - `id` - SAUCE comment block identification, this should be "COMNT"
- - `count`- The reported number of lines in the SAUCE comment block.
- - `lines` - Lines of text, each line should comprise of 64 characters.

---

### Similar projects and languages

- Go, [textmodes sauce](https://github.com/textmodes/sauce)
- Python, [Parser for SAUCE](https://pypi.org/project/sauce/)
- Elixir, [Saucexages](https://hexdocs.pm/saucexages/overview.html)
