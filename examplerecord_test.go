package sauce_test

import (
	"fmt"
	"log"

	"github.com/bengarrett/sauce"
)

func ExampleContains() {
	b, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		log.Fatal(err)
	}
	ok := sauce.Contains(b)
	fmt.Printf("File contains SAUCE metadata? %v", ok)
	// Output: File contains SAUCE metadata? true
}

func ExampleDecode() {
	b, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		log.Fatal(err)
	}
	sr := sauce.Decode(b)
	fmt.Printf("%+v", sr)
	// Output: {ID:SAUCE Version:00 Title:Sauce title Author:Sauce author Group:Sauce group Date:{Value:20161126 Time:2016-11-26 00:00:00 +0000 UTC Epoch:1480118400} FileSize:{Bytes:3713 Decimal:3.7 kB Binary:3.6 KiB} Data:{Type:text or character stream Name:text or character stream} File:{Type:0 Name:ASCII text} Info:{Info1:{Value:977 Info:character width} Info2:{Value:9 Info:number of lines} Info3:{Value:0 Info:} Flags:{Decimal:19 Binary:10011 B:{Flag:non-blink mode Info:non-blink mode} LS:{Flag:no preference Info:no preference} AR:{Flag:invalid value Info:invalid value} Interpretations:} Font:IBM VGA} Desc:ASCII text file with no formatting codes or color codes. Comnt:{ID:COMNT Count:1 Comment:[Any comments go here.                                           ]}}
}

func ExampleDecode_none() {
	b := []byte("This string of text does not contain any SAUCE.")
	sr := sauce.Decode(b)
	fmt.Printf("%+v", sr)
	// Output: {ID: Version: Title: Author: Group: Date:{Value: Time:0001-01-01 00:00:00 +0000 UTC Epoch:0} FileSize:{Bytes:0 Decimal: Binary:} Data:{Type:undefined Name:} File:{Type:0 Name:} Info:{Info1:{Value:0 Info:} Info2:{Value:0 Info:} Info3:{Value:0 Info:} Flags:{Decimal:0 Binary: B:{Flag:invalid value Info:} LS:{Flag:invalid value Info:} AR:{Flag:invalid value Info:} Interpretations:} Font:} Desc: Comnt:{ID: Count:0 Comment:[]}}
}

func ExampleIndex() {
	b, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		log.Fatal(err)
	}
	i := sauce.Index(b)
	fmt.Printf("SAUCE metadata position index: %v", i)
	// Output: SAUCE metadata position index: 1320
}

func ExampleJSON() {
	b, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		log.Fatal(err)
	}
	sr := sauce.Decode(b)
	js, err := sauce.JSON(&sr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(js))
	// Output: {"id":"SAUCE","version":"00","title":"Sauce title","author":"Sauce author","group":"Sauce group","date":{"value":"20161126","iso":"2016-11-26T00:00:00Z","epoch":1480118400},"filesize":{"bytes":3713,"decimal":"3.7 kB","binary":"3.6 KiB"},"dataType":{"type":1,"name":"text or character stream"},"fileType":{"type":0,"name":"ASCII text"},"typeInfo":{"1":{"value":977,"info":"character width"},"2":{"value":9,"info":"number of lines"},"3":{"value":0,"info":""},"flags":{"decimal":19,"binary":"10011","nonBlinkMode":{"flag":"1","interpretation":"non-blink mode"},"letterSpacing":{"flag":"00","interpretation":"no preference"},"aspectRatio":{"flag":"11","interpretation":"invalid value"}},"fontName":"IBM VGA"},"comments":{"id":"COMNT","count":1,"lines":["Any comments go here.                                           "]}}
}

func ExampleNewRecord() {
	file, err := static.Open("static/sauce.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sr, err := sauce.NewRecord(file)
	if err != nil {
		log.Fatal(err)
	}

	js, err := sauce.JSON(sr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(js))
	// Output: {"id":"SAUCE","version":"00","title":"Sauce title","author":"Sauce author","group":"Sauce group","date":{"value":"20161126","iso":"2016-11-26T00:00:00Z","epoch":1480118400},"filesize":{"bytes":3713,"decimal":"3.7 kB","binary":"3.6 KiB"},"dataType":{"type":1,"name":"text or character stream"},"fileType":{"type":0,"name":"ASCII text"},"typeInfo":{"1":{"value":977,"info":"character width"},"2":{"value":9,"info":"number of lines"},"3":{"value":0,"info":""},"flags":{"decimal":19,"binary":"10011","nonBlinkMode":{"flag":"1","interpretation":"non-blink mode"},"letterSpacing":{"flag":"00","interpretation":"no preference"},"aspectRatio":{"flag":"11","interpretation":"invalid value"}},"fontName":"IBM VGA"},"comments":{"id":"COMNT","count":1,"lines":["Any comments go here.                                           "]}}
}

func ExampleJSONIndent() {
	b, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		log.Fatal(err)
	}

	sr := sauce.Decode(b)

	const indent = "  "
	js, err := sauce.JSONIndent(&sr, indent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(js))
	// Output: {
	//   "id": "SAUCE",
	//   "version": "00",
	//   "title": "Sauce title",
	//   "author": "Sauce author",
	//   "group": "Sauce group",
	//   "date": {
	//     "value": "20161126",
	//     "iso": "2016-11-26T00:00:00Z",
	//     "epoch": 1480118400
	//   },
	//   "filesize": {
	//     "bytes": 3713,
	//     "decimal": "3.7 kB",
	//     "binary": "3.6 KiB"
	//   },
	//   "dataType": {
	//     "type": 1,
	//     "name": "text or character stream"
	//   },
	//   "fileType": {
	//     "type": 0,
	//     "name": "ASCII text"
	//   },
	//   "typeInfo": {
	//     "1": {
	//       "value": 977,
	//       "info": "character width"
	//     },
	//     "2": {
	//       "value": 9,
	//       "info": "number of lines"
	//     },
	//     "3": {
	//       "value": 0,
	//       "info": ""
	//     },
	//     "flags": {
	//       "decimal": 19,
	//       "binary": "10011",
	//       "nonBlinkMode": {
	//         "flag": "1",
	//         "interpretation": "non-blink mode"
	//       },
	//       "letterSpacing": {
	//         "flag": "00",
	//         "interpretation": "no preference"
	//       },
	//       "aspectRatio": {
	//         "flag": "11",
	//         "interpretation": "invalid value"
	//       }
	//     },
	//     "fontName": "IBM VGA"
	//   },
	//   "comments": {
	//     "id": "COMNT",
	//     "count": 1,
	//     "lines": [
	//       "Any comments go here.                                           "
	//     ]
	//   }
	// }
}