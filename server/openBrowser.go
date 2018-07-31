package main

import (
	"log"
	"net/http"
	"time"

	"github.com/skratchdot/open-golang/open"
)

func openBrowser(port string) {
	for {
		time.Sleep(time.Second)

		log.Println("Checking if started...")
		resp, err := http.Get("http://localhost" + port)
		if err != nil {
			log.Println("Failed:", err)
			continue
		}
		resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Println("Not OK:", resp.StatusCode)
			continue
		}

		// Reached this point: server is up and running!
		break
	}
	open.Run("http://localhost" + port)
}
