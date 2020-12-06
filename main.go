package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
)

var (
	conn redis.Conn
	err error
)

func handler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Incrementando accesos...")
	err = conn.Send("HINCRBY", "visits", "count", 1)
	if err != nil {
		log.Fatal("Error incrementing count", err)
	}

	hits, err := redis.String(conn.Do("HGET", "visits", "count"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, hits)
}

func main() {

	server := os.Getenv("REDIS_SERVER")
	if len(server) == 0 {
		server = "localhost"
	}

	conn, err = redis.Dial("tcp", server + ":6379")
	if err != nil {
		log.Fatal("Error conectando con redis en: " + server + "\n", err)
	}
	defer conn.Close()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
