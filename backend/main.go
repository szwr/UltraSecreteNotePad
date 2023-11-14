package main

import (
	"flag"
	"fmt"

	"example.com/api"
	"example.com/db"
)

func main() {
	listenAddr := flag.String("l", ":3000", "Server Address")
	databaseAddr := flag.String("db", ":6379", "Database Address")
	databasePass := flag.String("dbpass", "", "Database Password")
	flag.Parse()

	rc := db.ReturnRedisClient(*databaseAddr, *databasePass, 0)
	server := api.NewServer(*listenAddr, rc)

	fmt.Printf("Starting Server: %v\n", *listenAddr)
	server.Start()
}
