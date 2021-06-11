package config

import (
	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

const AllowedCORSDomain = "http://localhost"
