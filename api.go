package main

import (
	"log"

	"github.com/howardliam/music-tab-api/config"
)

func main() {
	conf := config.LoadConfig()
	log.Printf("Config: %v", conf)
}
