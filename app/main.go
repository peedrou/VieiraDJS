package main

import (
	// crud "VieiraDJS/app/db/CRUD"
	// "VieiraDJS/app/services/jobs"
	// "VieiraDJS/app/services/users"
	// "fmt"
	"VieiraDJS/app/kafka"
	"log"
	"os"
	"strconv"

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

	producer, err := kafka.NewKafkaProducer([]string{kafkaBrokers})

	producer.SendMessage("my_new_topic", "Hello Im a message!")

	producer.Close()

	// session, err := cluster.CreateSession()
	// if err != nil {
	// 	log.Fatalf("Failed to connect to Cassandra: %v", err)
	// }
	// defer session.Close()

	// user_id, err := users.RegisterUser(session, "testuser6", "hellopassword", "thisisanemail6@gmail.com")
	// if err != nil {
	// 	fmt.Printf("Error creating user: %v\n", err)
	// 	return
	// }

	// err = jobs.CreateJob(session, user_id, true, 3, time.Now(), "8H")
	// if err != nil {
	// 	fmt.Printf("Error creating job: %v\n", err)
	// 	return
	// }

	// fmt.Println("Job successfully created and inserted into Cassandra!")

	// result, _ := crud.ReadModel(session, "jobs", []string{"job_id"}, []string{"interval"}, "2h")

	// fmt.Printf("job successfully read from Cassandra! %v", result)

	// err = crud.UpdateModelBatch(session, "jobs", "interval", "6H", "job_id", result)
	// if err != nil {
	// 	fmt.Printf("Error Deleting model: %v\n", err)
	// 	return
	// }

	// err = crud.RemoveModel(session, "jobs", "job_id", result)
	// if err != nil {
	// 	fmt.Printf("Error Deleting model: %v\n", err)
	// 	return
	// }
}

func parsePort(port string) int {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}
	return portInt
}
