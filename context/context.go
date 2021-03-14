package context

import (
	"GOIF/types"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Start(properties types.Properties) {
	file, err := ioutil.ReadFile("banner.txt")
	if err != nil {
		log.Fatal("Fail to load GOIF banner logo.")
	}
	log.Print(string(file))
	log.Print("Starting GOIF context...")

	log.Print("Loading routes...")
	properties.RegisterRoutes()
	log.Print("Routes started.")

	address := fmt.Sprintf("%s:%d", properties.Host, properties.Port)

	log.Printf("Server configured for running at: %s", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal(err)
	}
}
