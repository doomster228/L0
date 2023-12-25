package common

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	stan "github.com/nats-io/stan.go"
)

// Подключение к кластеру NATS Streaming
func NatsConnection(clientID string) stan.Conn {
	sc, err := stan.Connect("test-cluster", clientID, stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS Streaming: %v", err)
	}
	return sc
}

// Открытие соединения с базой данных
func DatabaseConnection() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=1 dbname=test0 sslmode=disable")
	if err != nil {
		log.Fatalf("Ошибка подключения к PostgreSQL: %v", err)
	}
	return db
}
