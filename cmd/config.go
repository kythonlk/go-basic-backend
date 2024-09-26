package cmd

import (
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

var (
	JwtSecret        string
	RateLimitShort   time.Duration
	RateLimitDaily   int
	ConnectionString string
)

func Init() {

	JwtSecret = getEnv("JWT_TOKEN", "")
	ConnectionString = getEnv("SQL_STRING", "")
	RateLimitShort = getDuration("RATE_LIMIT_SHORT", 30*time.Second)
	RateLimitDaily = getInt("RATE_LIMIT_DAILY", 5)

}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getInt(key string, fallback int) int {
	strValue := getEnv(key, "")
	if value, err := strconv.Atoi(strValue); err == nil {
		return value
	}
	return fallback
}

func getDuration(key string, fallback time.Duration) time.Duration {
	strValue := getEnv(key, "")
	if value, err := time.ParseDuration(strValue); err == nil {
		return value
	}
	return fallback
}
