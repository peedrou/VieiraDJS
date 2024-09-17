package session

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

type SessionCreator interface {
    CreateSession() (*gocql.Session, error)
}

type RealSessionCreator struct{}

func init() {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func CreateSession() (*gocql.Session, error) {
	cluster := gocql.NewCluster(os.Getenv("CASSANDRA_HOSTS"))
    cluster.Port = getEnvAsInt("CASSANDRA_PORT", 9042)
    cluster.Keyspace = os.Getenv("CASSANDRA_KEYSPACE")
    cluster.Authenticator = gocql.PasswordAuthenticator{
        Username: os.Getenv("CASSANDRA_USERNAME"),
        Password: os.Getenv("CASSANDRA_PASSWORD"),
    }
    cluster.Consistency = gocql.Quorum

    session, err := cluster.CreateSession()
    if err != nil {
        return nil, fmt.Errorf("unable to create Cassandra session: %w", err)
    }

    return session, nil
}

func getEnvAsInt(name string, defaultValue int) int {
    if valueStr := os.Getenv(name); valueStr != "" {
        if value, err := strconv.Atoi(valueStr); err == nil {
            return value
        }
    }
    return defaultValue
}