package main

import (
	// crud "VieiraDJS/app/db/CRUD"
	// "VieiraDJS/app/services/jobs"
	// "VieiraDJS/app/services/users"
	"VieiraDJS/app/kafka"
	"VieiraDJS/app/services/jobs"
	"VieiraDJS/app/services/scheduler"
	"VieiraDJS/app/services/users"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	// "time"

	"strings"

	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cassandraHosts := os.Getenv("CASSANDRA_HOSTS")
	cassandraPort := os.Getenv("CASSANDRA_PORT")
	cassandraKeyspace := os.Getenv("CASSANDRA_KEYSPACE")
	kafkaBrokers := os.Getenv("KAFKA_BROKERS")

	cluster := gocql.NewCluster(strings.Split(cassandraHosts, ",")...)
	cluster.Port = parsePort(cassandraPort)
	cluster.Keyspace = cassandraKeyspace
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to Cassandra: %v", err)
	}
	defer session.Close()

	user_id, err := users.RegisterUser(session, "testuser6", "hellopassword", "thisisanemail6@gmail.com")
	if err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		return
	}

	err = jobs.CreateJob(session, user_id, true, 3, time.Now(), "8H")
	if err != nil {
		fmt.Printf("Error creating job: %v\n", err)
		return
	}

	fmt.Println("Job successfully created and inserted into Cassandra!")

	producer, err := kafka.NewKafkaProducer([]string{kafkaBrokers})

	if err != nil {
		fmt.Printf("Error creating producer: %v\n", err)
		return
	}

	tasks, err := scheduler.CheckPendingTasks(session)

	if err != nil {
		fmt.Printf("Error checking pending tasks: %v\n", err)
		return
	}

	if tasks != nil {
		tasksSucceeded, tasksFailed, err := scheduler.SchedulePendingTasks(producer, tasks, session)
		fmt.Printf("%s , %s , %s", tasksSucceeded, tasksFailed, err)
	}

	producer.Close()

}

func parsePort(port string) int {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}
	return portInt
}
