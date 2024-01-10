package sauce_test

import (
	"bytes"
	"fmt"
	"log"

	"github.com/bengarrett/sauce"
)

func ExampleContains() {
	b, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(sauce.Contains(b))
	// Output: true
}

func ExampleContains_none() {
	b := []byte("This string of text does not contain any SAUCE.")
	fmt.Print(sauce.Contains(b))
	// Output: false
}

func ExampleIndex() {
	b, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		fmt.Print(err)
	}
	i := sauce.Index(b)
	fmt.Printf("SAUCE metadata position index: %v", i)
	// Output: SAUCE metadata position index: 1190
}

func ExampleTrim() {
	b, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		fmt.Print(err)
	}

	t := sauce.Trim(b)

	fmt.Printf("The original size of the file is %d bytes and contains sauce? %v\n", len(b), sauce.Contains(b))
	fmt.Printf("The trimmed size of the file is %d bytes and contains sauce? %v\n", len(t), sauce.Contains(t))
	// Output: The original size of the file is 1318 bytes and contains sauce? true
	// The trimmed size of the file is 1119 bytes and contains sauce? false
}

func ExampleTrim_comnt() {
	b, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		fmt.Print(err)
	}

	// normalize line endings
	b = bytes.ReplaceAll(b, []byte("\r\n"), []byte("\n"))

	// print the raw sauce binary data
	const index = 7
	if x := bytes.Split(b, []byte("\n")); len(x) >= index+1 {
		fmt.Printf("\nSAUCE: %q\n\n", x[index])
	}

	// print the trimmed text
	fmt.Printf("%s", string(sauce.Trim(b)))
	// Output: SAUCE: "\x1aCOMNTAny comments go here.                                           SAUCE00Sauce title                        Sauce author        Sauce group         20161126\x9d\x0e\x00\x00\x01\x00\xd1\x03\t\x00\x00\x00\x00\x00\x01\x13IBM VGA\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
	//
	// File: static/sauce.txt
	//
	// This must be opened and saved as a Windows 1252 document.
	//
	// SAUCE package test with complete data and comment.
	//
	// Cras sit amet purus urna. Phasellus in dapibus ex. Proin pretium eget leo ut gravida. Praesent egestas urna at tincidunt mollis. Vivamus nec urna lorem. Vestibulum molestie accumsan lectus, in egestas metus facilisis eget. Nam consectetur, metus et sodales aliquam, mi dui dapibus metus, non cursus libero felis ac nunc. Nulla euismod, turpis sed mollis faucibus, orci elit dapibus leo, vitae placerat diam eros sed velit. Fusce convallis, lorem ut vulputate suscipit, tortor risus rhoncus mauris, a mattis metus magna at lorem. Sed molestie velit ipsum, in vulputate metus consequat eget. Fusce quis dui lacinia, laoreet lectus et, luctus velit. Pellentesque ut nisi quis orci pulvinar placerat vel ac lorem. Maecenas finibus fermentum erat, a pulvinar augue dictum mattis. Aenean vulputate consectetur velit at dictum. Donec vehicula ante quis ante venenatis, eu ultrices lectus egestas. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae;
}

func ExampleTrim_nocomnt() {
	b, err := static.ReadFile("static/sauce-nocomnt.txt")
	if err != nil {
		fmt.Print(err)
	}

	// normalize line endings
	b = bytes.ReplaceAll(b, []byte("\r\n"), []byte("\n"))

	// print the raw sauce binary data
	const index = 7
	if x := bytes.Split(b, []byte("\n")); len(x) >= index+1 {
		fmt.Printf("\nSAUCE: %q\n\n", x[index])
	}

	// print the trimmed text
	fmt.Printf("%s", string(sauce.Trim(b)))
	// Output: SAUCE: "\x1aSAUCE00Sauce title                        Sauce author        Sauce group         20161126\x9d\x0e\x00\x00\x01\x00\xd1\x03\t\x00\x00\x00\x00\x00\x01\x13IBM VGA\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
	//
	// File: static/sauce-nocomnt.txt
	//
	// This must be opened and saved as a Windows 1252 document.
	//
	// SAUCE package test with complete data and comment.
	//
	// Cras sit amet purus urna. Phasellus in dapibus ex. Proin pretium eget leo ut gravida. Praesent egestas urna at tincidunt mollis. Vivamus nec urna lorem. Vestibulum molestie accumsan lectus, in egestas metus facilisis eget. Nam consectetur, metus et sodales aliquam, mi dui dapibus metus, non cursus libero felis ac nunc. Nulla euismod, turpis sed mollis faucibus, orci elit dapibus leo, vitae placerat diam eros sed velit. Fusce convallis, lorem ut vulputate suscipit, tortor risus rhoncus mauris, a mattis metus magna at lorem. Sed molestie velit ipsum, in vulputate metus consequat eget. Fusce quis dui lacinia, laoreet lectus et, luctus velit. Pellentesque ut nisi quis orci pulvinar placerat vel ac lorem. Maecenas finibus fermentum erat, a pulvinar augue dictum mattis. Aenean vulputate consectetur velit at dictum. Donec vehicula ante quis ante venenatis, eu ultrices lectus egestas. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae;
}

//nolint:lll
func ExampleDecode() {
	b, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		fmt.Print(err)
	}

	sr := sauce.Decode(b)
	fmt.Printf("%+v", sr)
	// Output: {ID:SAUCE Version:00 Title:Sauce title Author:Sauce author Group:Sauce group Date:{Value:20161126 Time:2016-11-26 00:00:00 +0000 UTC Epoch:1480118400} FileSize:{Bytes:3741 Decimal:3.7 kB Binary:3.7 KiB} Data:{Type:text or character stream Name:text or character stream} File:{Type:0 Name:ASCII text} Info:{Info1:{Value:977 Info:character width} Info2:{Value:9 Info:number of lines} Info3:{Value:0 Info:} Flags:{Decimal:19 Binary:10011 B:{Flag:non-blink mode Info:non-blink mode} LS:{Flag:no preference Info:no preference} AR:{Flag:invalid value Info:invalid value} Interpretations:} Font:IBM VGA} Desc:ASCII text file with no formatting codes or color codes. Comnt:{ID:COMNT Count:1 Index:1121 Comment:[Any comments go here.                                           ]}}
}

//nolint:lll
func ExampleDecode_none() {
	b := []byte("This string of text does not contain any SAUCE.")

	sr := sauce.Decode(b)
	fmt.Printf("%+v", sr)
	// Output: {ID: Version: Title: Author: Group: Date:{Value: Time:0001-01-01 00:00:00 +0000 UTC Epoch:0} FileSize:{Bytes:0 Decimal: Binary:} Data:{Type:undefined Name:} File:{Type:0 Name:} Info:{Info1:{Value:0 Info:} Info2:{Value:0 Info:} Info3:{Value:0 Info:} Flags:{Decimal:0 Binary: B:{Flag:invalid value Info:} LS:{Flag:invalid value Info:} AR:{Flag:invalid value Info:} Interpretations:} Font:} Desc: Comnt:{ID: Count:0 Index:-1 Comment:[]}}
}

//nolint:lll
func ExampleRead() {
	// open file
	file, err := static.Open("static/sauce.txt")
	if err != nil {
		log.Print(err)
		return
	}
	defer file.Close()

	// create a new sauce record
	sr, err := sauce.Read(file)
	if err != nil {
		log.Print(err)
		return
	}

	// encode and print the sauce record as JSON
	js, err := sr.JSON()
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Print(string(js))
	// Output: {"id":"SAUCE","version":"00","title":"Sauce title","author":"Sauce author","group":"Sauce group","date":{"value":"20161126","iso":"2016-11-26T00:00:00Z","epoch":1480118400},"filesize":{"bytes":3741,"decimal":"3.7 kB","binary":"3.7 KiB"},"dataType":{"type":1,"name":"text or character stream"},"fileType":{"type":0,"name":"ASCII text"},"typeInfo":{"1":{"value":977,"info":"character width"},"2":{"value":9,"info":"number of lines"},"3":{"value":0,"info":""},"flags":{"decimal":19,"binary":"10011","nonBlinkMode":{"flag":"1","interpretation":"non-blink mode"},"letterSpacing":{"flag":"00","interpretation":"no preference"},"aspectRatio":{"flag":"11","interpretation":"invalid value"}},"fontName":"IBM VGA"},"comments":{"id":"COMNT","count":1,"lines":["Any comments go here.                                           "]}}
}

func ExampleRead_none() {
	b := []byte("This string of text does not contain any SAUCE.")

	// create a new sauce record
	sr, err := sauce.Read(bytes.NewReader(b))
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	// encode and print the sauce record as JSON
	js, err := sr.JSON()
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Print(string(js))
	// Output: {"id":"","version":"","title":"","author":"","group":"","date":{"value":"","iso":"0001-01-01T00:00:00Z","epoch":0},"filesize":{"bytes":0,"decimal":"","binary":""},"dataType":{"type":0,"name":""},"fileType":{"type":0,"name":""},"typeInfo":{"1":{"value":0,"info":""},"2":{"value":0,"info":""},"3":{"value":0,"info":""},"flags":{"decimal":0,"binary":"","nonBlinkMode":{"flag":"","interpretation":""},"letterSpacing":{"flag":"","interpretation":""},"aspectRatio":{"flag":"","interpretation":""}},"fontName":""},"comments":{"id":"","count":0,"lines":[]}}
}

//nolint:lll
func ExampleRecord_JSON() {
	b, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		fmt.Print(err)
	}

	sr := sauce.Decode(b)
	js, err := sr.JSON()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(js))
	// Output: {"id":"SAUCE","version":"00","title":"Sauce title","author":"Sauce author","group":"Sauce group","date":{"value":"20161126","iso":"2016-11-26T00:00:00Z","epoch":1480118400},"filesize":{"bytes":3741,"decimal":"3.7 kB","binary":"3.7 KiB"},"dataType":{"type":1,"name":"text or character stream"},"fileType":{"type":0,"name":"ASCII text"},"typeInfo":{"1":{"value":977,"info":"character width"},"2":{"value":9,"info":"number of lines"},"3":{"value":0,"info":""},"flags":{"decimal":19,"binary":"10011","nonBlinkMode":{"flag":"1","interpretation":"non-blink mode"},"letterSpacing":{"flag":"00","interpretation":"no preference"},"aspectRatio":{"flag":"11","interpretation":"invalid value"}},"fontName":"IBM VGA"},"comments":{"id":"COMNT","count":1,"lines":["Any comments go here.                                           "]}}
}

func ExampleRecord_JSONIndent() { //nolint:funlen
	b, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		fmt.Print(err)
	}

	sr := sauce.Decode(b)

	const indent = "  "
	js, err := sr.JSONIndent(indent)
	if err != nil {
		fmt.Print(err)
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
	//     "bytes": 3741,
	//     "decimal": "3.7 kB",
	//     "binary": "3.7 KiB"
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

func ExampleRecord_Valid() {
	b, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		fmt.Print(err)
	}

	sr := sauce.Decode(b)
	fmt.Print(sr.Valid())
	// Output: true
}

func ExampleRecord_Valid_invalid() {
	b := []byte("This string of text does not contain any SAUCE.")
	sr := sauce.Decode(b)

	fmt.Print(sr.Valid())
	// Output: false
}

//nolint:lll
func ExampleRecord_XML() {
	b, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		fmt.Print(err)
	}
	sr := sauce.Decode(b)
	xm, err := sr.XML()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(xm))
	// Output: <Record id="SAUCE" version="00"><title>Sauce title</title><author>Sauce author</author><group>Sauce group</group><date epoch="1480118400"><value>20161126</value><date>2016-11-26T00:00:00Z</date></date><filesize decimal="3.7 kB" binary="3.7 KiB"><bytes>3741</bytes></filesize><data_type><type>1</type><name>text or character stream</name></data_type><file_type><type>0</type><name>ASCII text</name></file_type><type_info><type1 type="character width"><value>977</value></type1><type2 type="number of lines"><value>9</value></type2><type3 type=""><value>0</value></type3><flags decimal="19" binary="10011"><non_blink_mode interpretation="non-blink mode"><flag>1</flag></non_blink_mode><letter_spacing interpretation="no preference"><flag>00</flag></letter_spacing><aspect_ratio interpretation="invalid value"><flag>11</flag></aspect_ratio></flags><fontname>IBM VGA</fontname></type_info><comments id="COMNT" count="1"><line>Any comments go here.                                           </line></comments></Record>
}

func ExampleRecord_XMLIndent() {
	b, err := static.ReadFile("static/sauce.txt")
	if err != nil {
		fmt.Print(err)
	}
	sr := sauce.Decode(b)
	const indent = "  "
	xm, err := sr.XMLIndent(indent)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(xm))
	// Output: <Record id="SAUCE" version="00">
	//   <title>Sauce title</title>
	//   <author>Sauce author</author>
	//   <group>Sauce group</group>
	//   <date epoch="1480118400">
	//     <value>20161126</value>
	//     <date>2016-11-26T00:00:00Z</date>
	//   </date>
	//   <filesize decimal="3.7 kB" binary="3.7 KiB">
	//     <bytes>3741</bytes>
	//   </filesize>
	//   <data_type>
	//     <type>1</type>
	//     <name>text or character stream</name>
	//   </data_type>
	//   <file_type>
	//     <type>0</type>
	//     <name>ASCII text</name>
	//   </file_type>
	//   <type_info>
	//     <type1 type="character width">
	//       <value>977</value>
	//     </type1>
	//     <type2 type="number of lines">
	//       <value>9</value>
	//     </type2>
	//     <type3 type="">
	//       <value>0</value>
	//     </type3>
	//     <flags decimal="19" binary="10011">
	//       <non_blink_mode interpretation="non-blink mode">
	//         <flag>1</flag>
	//       </non_blink_mode>
	//       <letter_spacing interpretation="no preference">
	//         <flag>00</flag>
	//       </letter_spacing>
	//       <aspect_ratio interpretation="invalid value">
	//         <flag>11</flag>
	//       </aspect_ratio>
	//     </flags>
	//     <fontname>IBM VGA</fontname>
	//   </type_info>
	//   <comments id="COMNT" count="1">
	//     <line>Any comments go here.                                           </line>
	//   </comments>
	// </Record>
}
