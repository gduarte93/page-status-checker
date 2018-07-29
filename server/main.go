package main

import (
	"fmt"
	"net/http"
)

const port = ":4000"

func main() {
	http.Handle("/", http.FileServer(http.Dir("../client")))
	http.HandleFunc("/getStatus", getStatusHandler)

	fmt.Println("Listening on http://localhost" + port)
	http.ListenAndServe(port, nil)
}
