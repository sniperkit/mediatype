package main

import (
	"fmt"

	mediatype "github.com/sniperkit/mediatype/pkg"
)

func main() {
	mutable := &mediatype.Mutable{
		Main: "application",
		Sub:  "json",
	}

	mutable.String() // "application/json"

	mutable.Sub = "xhtml"
	mutable.Suf = "xml"
	fmt.Println("mutable.String(): ", mutable.String()) // "application/xhtml+xml"

	immutable := mutable.Immutable()
	mutable.Main = "image"
	immutable.String()                                    // "application/xhtml+xml"
	fmt.Println("mutable.String(): ", immutable.String()) // "application/xhtml+xml"
}
