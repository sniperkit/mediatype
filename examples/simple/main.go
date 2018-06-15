package main

import (
	"fmt"

	mediatype "github.com/sniperkit/mediatype/pkg"
)

func main() {

	contentTypeString := "application/vnd.google-earth.kml+xml; charset=utf-8"
	mediaType, _ := mediatype.Parse(contentTypeString)

	fmt.Println("MainType: ", mediaType.MainType())     // "application"
	fmt.Println("SubType: ", mediaType.SubType())       // "kml"
	fmt.Println("Trees: ", mediaType.Trees())           // ["vnd", "google-earth"]
	fmt.Println("Prefix: ", mediaType.Prefix())         // "vnd"
	fmt.Println("Suffix: ", mediaType.Suffix())         // "xml"
	fmt.Println("Parameters: ", mediaType.Parameters()) // ["charset": "utf-8"]

	fmt.Println("FullType: ", mediaType.FullType()) // "application/vnd.google-earth.kml+xml"
	fmt.Println("String: ", mediaType.String())     // "application/vnd.google-earth.kml+xml; charset=utf-8"

}
