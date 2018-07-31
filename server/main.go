package main

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/packr"
)

const port = ":4000"

func main() {
	box := packr.NewBox("../client")

	// use http.Dir("../client") instead of box when you don't want to package with packr
	//
	// make sure to update Makefile location/build,
	// since not using packr makes binary location important:
	// must be in the server dir to recognize html/css/js files
	http.Handle("/", http.FileServer(box))
	http.HandleFunc("/getStatus", getStatusHandler)

	go openBrowser(port)

	fmt.Println("Listening on http://localhost" + port)
	http.ListenAndServe(port, nil)
}
