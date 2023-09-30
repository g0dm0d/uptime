package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const version = "BUILD_VERSION"

type App struct {
	Addr    string
	Port    int
	Secret  string
	Version string
}

type InfluxDB struct {
	Addr   string
	Port   int
	Org    string
	Bucket string
	Token  string
}

type PostgresDB struct {
	DNS string
}

type Config struct {
	App        App
	InfluxDB   InfluxDB
	PostgresDB PostgresDB
}

// New returns a new Config struct
func New() *Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	return &Config{
		App: App{
			Addr:    getEnv("APP_ADDR", "localhost"),
			Port:    getEnvAsInt("APP_PORT", 3000),
			Secret:  getEnv("APP_AUTH_SECRET", "secret"),
			Version: version,
		},
		InfluxDB: InfluxDB{
			Addr:   getEnv("INFLUXDB_ADDR", "http://localhost"),
			Port:   getEnvAsInt("INFLUXDB_PORT", 8086),
			Org:    getEnv("INFLUXDB_ORG", "uptime"),
			Bucket: getEnv("INFLUXDB_BUCKET", "uptime"),
			Token:  getEnv("INFLUXDB_ADMIN_TOKEN", "token"),
		},
		PostgresDB: PostgresDB{
			DNS: getEnv("POSTGRES_DSN", "postgres://postgres:12345@localhost:5432/postgres?sslmode=disable"),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
