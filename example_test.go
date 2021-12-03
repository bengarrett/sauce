package sauce_test

import (
	"embed"
	"fmt"
	"log"
	"time"

	"github.com/bengarrett/sauce"
)

//go:embed static/*
var static embed.FS

func Example() {
	// open a file
	file, err := static.Open("static/sauce.txt")
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
	if !sr.Valid() {
		fmt.Println("no sauce metadata found")
		return
	}

	// print specific SAUCE fields
	fmt.Printf("%q\n", sr.Title)
	fmt.Printf("Author:\t%s.\n", sr.Author)
	fmt.Printf("Group:\t%s.\n", sr.Group)
	fmt.Printf("Date:\t%s.\n", sr.Date.Time.Format(time.ANSIC))

	// return the SAUCE data as indented JSON
	js, err := sr.JSONIndent("    ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%s", js)

	// Output: "Sauce title"
	// Author:	Sauce author.
	// Group:	Sauce group.
	// Date:	Sat Nov 26 00:00:00 2016.
	//{
	//     "id": "SAUCE",
	//     "version": "00",
	//     "title": "Sauce title",
	//     "author": "Sauce author",
	//     "group": "Sauce group",
	//     "date": {
	//         "value": "20161126",
	//         "iso": "2016-11-26T00:00:00Z",
	//         "epoch": 1480118400
	//     },
	//     "filesize": {
	//         "bytes": 3713,
	//         "decimal": "3.7 kB",
	//         "binary": "3.6 KiB"
	//     },
	//     "dataType": {
	//         "type": 1,
	//         "name": "text or character stream"
	//     },
	//     "fileType": {
	//         "type": 0,
	//         "name": "ASCII text"
	//     },
	//     "typeInfo": {
	//         "1": {
	//             "value": 977,
	//             "info": "character width"
	//         },
	//         "2": {
	//             "value": 9,
	//             "info": "number of lines"
	//         },
	//         "3": {
	//             "value": 0,
	//             "info": ""
	//         },
	//         "flags": {
	//             "decimal": 19,
	//             "binary": "10011",
	//             "nonBlinkMode": {
	//                 "flag": "1",
	//                 "interpretation": "non-blink mode"
	//             },
	//             "letterSpacing": {
	//                 "flag": "00",
	//                 "interpretation": "no preference"
	//             },
	//             "aspectRatio": {
	//                 "flag": "11",
	//                 "interpretation": "invalid value"
	//             }
	//         },
	//         "fontName": "IBM VGA"
	//     },
	//     "comments": {
	//         "id": "COMNT",
	//         "count": 1,
	//         "lines": [
	//             "Any comments go here.                                           "
	//         ]
	//     }
	// }
}
