package global

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

const (
	DefaultURLPathDB = "postgresql://postgres:postgres@localhost:5432/d_track"
	DefaultVerApp    = "0.1.0"
	DefaultBuildApp  = "20230302.1"
	DefaultSrvAddr   = "localhost:8080"
)

type ServiceConfig struct {
	URLPathDB string
	IsLog     bool
	Ver       string
	Build     string
	SrvAddr   string
}

var Config ServiceConfig

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found err:", err)
	}

	Config = ServiceConfig{
		URLPathDB: getEnv("URL_PATH_DB", DefaultURLPathDB),
		IsLog:     getEnvAsBool("IS_LOG", false),
		Ver:       getEnv("VER", DefaultVerApp),
		Build:     getEnv("BUILD", DefaultBuildApp),
		SrvAddr:   getEnv("SRV_ADDR", DefaultSrvAddr),
	}
}

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
