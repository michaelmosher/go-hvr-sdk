package main

import (
	"fmt"
	"log"

	"github.com/michaelmosher/go-hvr-sdk/hvrhub"
	hvrhub_postgresql "github.com/michaelmosher/go-hvr-sdk/hvrhub/postgresql"
)

func main() {
	connStr := "host=localhost user=michaelmosher dbname=hvrhub sslmode=disable"

	hub_client, err := hvrhub_postgresql.New(connStr)
	if err != nil {
		log.Fatal(err)
	}

	hub := hvrhub.Service{Client: hub_client}

	name := "hello"

	if err = hub.GetLocation(hvrhub.Location{Name: name}); err != nil {
		log.Fatal(err)
	}

	location, err := hub.GetLocation(name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hey, look at my record: %+v\n", location)
}