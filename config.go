package main

import "os"

type Config struct {
	ClientID     string
	ClientSecret string
}

func LoadConfig() Config {
	return Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}
}
