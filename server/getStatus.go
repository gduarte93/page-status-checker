package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"
)

func getStatusHandler(w http.ResponseWriter, r *http.Request) {
	pages := make(map[string]map[string]int)
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://laksjdfhaksjlhdf.com",
		"http://imgur.com",
	}

	c := make(chan []string)

	for _, link := range links {
		go checkLink(link, c)
	}

	i := 0
	for l := range c {
		num, _ := strconv.Atoi(l[1])
		if len(l) > 2 {
			code, _ := strconv.Atoi(l[2])
			pages[l[0]] = map[string]int{"status": num, "code": code}
		} else {
			pages[l[0]] = map[string]int{"status": num}
		}

		if i >= len(links)-1 {
			close(c)
		}
		i++
	}

	jPages, err := json.Marshal(pages)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Fprintf(w, string(jPages))
}

func checkLink(link string, c chan []string) {
	// If page takes longer than 10 seconds to respond, timeout
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(link)
	if e, ok := err.(net.Error); ok && e.Timeout() {
		fmt.Println("Error:", err)
		c <- []string{link, "2"} // 2 when page takes too long to respond
		return
	} else if err != nil {
		fmt.Println("Error:", err)
		c <- []string{link, "0"} // 0 when page is down
		return
	}

	fmt.Println(link + " looks up: " + strconv.Itoa(resp.StatusCode))
	if resp.StatusCode == 200 {
		c <- []string{link, "1", strconv.Itoa(resp.StatusCode)} // 1 when page is serving 200
	} else {
		c <- []string{link, "2", strconv.Itoa(resp.StatusCode)} // 2 when page is serving anything, but 200
	}

}
