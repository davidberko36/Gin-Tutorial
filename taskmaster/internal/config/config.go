package config


import (
	"log"
	"os"
)

type Config struct {
	DBURL string
	Port string
}

func Load() Config {
	dbURL := os.Getenv("DATABASE_URL") //fetches from environment variables
	port := os.Getenv("PORT")
	if dbURL == "" || port == "" {
		log.Fatal("Missing required environment variables")
	}
	return Config{DBURL: dbURL, Port: port}
}