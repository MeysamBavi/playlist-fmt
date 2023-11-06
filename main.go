package main

import (
	"context"
	"flag"
	"github.com/zmb3/spotify/v2"
	"log"
	"os"
)

var playlistId = flag.String("id", "", "The playlist id")

func main() {
	flag.Parse()

	cfg := LoadConfig()

	template, err := getTemplate()
	if err != nil {
		log.Fatalf("could not load template: %v\n", err)
	}

	client, err := createClient(cfg)
	if err != nil {
		log.Fatalf("could create client: %v\n", err)
	}

	playlist, err := client.GetPlaylist(context.Background(), spotify.ID(*playlistId))
	if err != nil {
		log.Fatalf("could not get playlist: %v\n", err)
	}

	shortPlaylist := NewPlaylist(playlist)
	err = template.Execute(os.Stdout, shortPlaylist)
	if err != nil {
		log.Fatalf("could not execute template: %v", err)
	}
}
