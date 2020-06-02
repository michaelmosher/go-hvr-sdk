package main

import (
	"fmt"
	"log"

	"github.com/michaelmosher/go-hvr-sdk/hvrhub"
	hvrhub_postgresql "github.com/michaelmosher/go-hvr-sdk/hvrhub/postgresql"
)

func main() {
	connStr := "host=localhost user=michaelmosher dbname=hvrhub sslmode=disable"

	pgClient, err := hvrhub_postgresql.New(connStr)
	if err != nil {
		log.Fatal(err)
	}

	hub := hvrhub.Service{Client: pgClient}

	name := "hello"

	if err := hub.NewChannel(hvrhub.Channel{"Test", "This is for testing"}); err != nil {
		log.Fatal(err)
	}

	if err := hub.UpdateChannel(hvrhub.Channel{"Test", "Test some more"}); err != nil {
		log.Fatal(err)
	}

	if chn, err := hub.GetChannel("Test"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Hey, look at my channel: %+v\n", chn)
	}

	if err := hub.DeleteChannel("Test"); err != nil {
		log.Fatal(err)
	}

	// if err = hub.DeleteLocation(name); err != nil {
	// 	log.Fatal(err)
	// }

	location, err := hub.GetLocation(name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hey, look at my record: %+v\n", location)
}
