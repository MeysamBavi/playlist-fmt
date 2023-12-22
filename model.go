package main

import (
	"github.com/zmb3/spotify/v2"
	"time"
)

type Track struct {
	Name     string
	Album    string
	Duration time.Duration
	Artists  []string
	URL      string
}

type Playlist struct {
	Name                  string
	Description           string
	URL                   string
	EstimatedCreationTime time.Time
	Tracks                []Track
}

func NewPlaylist(playlist *spotify.FullPlaylist) Playlist {
	p := Playlist{
		Name:        playlist.Name,
		Description: playlist.Description,
		URL:         getSpotifyURL(playlist.ExternalURLs),
		Tracks:      nil,
	}

	var averageTimeUnix int64
	count := 0
	for _, t := range playlist.Tracks.Tracks {
		track := Track{
			Name:     t.Track.Name,
			Album:    t.Track.Album.Name,
			Duration: t.Track.TimeDuration(),
			Artists:  nil,
			URL:      getSpotifyURL(t.Track.ExternalURLs),
		}
		for _, a := range t.Track.Artists {
			track.Artists = append(track.Artists, a.Name)
		}

		p.Tracks = append(p.Tracks, track)

		addedAt, err := time.Parse(time.RFC3339, t.AddedAt)
		if err == nil {
			averageTimeUnix += addedAt.Unix()
			count++
		}
	}
	if count > 0 {
		averageTimeUnix /= int64(count)
		p.EstimatedCreationTime = time.Unix(averageTimeUnix, 0)
	}

	return p
}

func getSpotifyURL(urls map[string]string) string {
	return urls["spotify"]
}
