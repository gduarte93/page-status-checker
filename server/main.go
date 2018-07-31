package main

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/packr"
)

const port = ":4000"

func main() {
	box := packr.NewBox("../client")

	http.Handle("/", http.FileServer(box))
	http.HandleFunc("/getStatus", getStatusHandler)

	go openBrowser(port)

	fmt.Println("Listening on http://localhost" + port)
	http.ListenAndServe(port, nil)
}
