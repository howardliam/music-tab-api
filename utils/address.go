package utils

import (
	"strconv"

	"github.com/howardliam/music-tab-api/config"
)

func GenerateAddress(server config.ServerConfig) string {
	return ":" + strconv.Itoa(int(server.Port))
}
