// Copyright (c) 2020 Ramon Quitales
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rquitales/go-memdb/memdb"
)

func main() {
	db := memdb.NewDB()

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("go-memdb CLI package example:\n---------------------\n\n")

	for {
		query := readStdIn(reader)
		if len(query) == 0 {
			continue
		}

		switch strings.ToLower(query[0]) {
		case "set":
			set(db, query)
		case "get":
			get(db, query)
		case "delete":
			deleteEntry(db, query)
		case "count":
			count(db, query)
		case "end":
			log.Println("Exiting go-memdb...")
			os.Exit(0)
		default:
			log.Println("[WARN] unsupported query")
		}
	}
}

func set(db *memdb.MemDB, query []string) {
	if len(query) != 3 {
		log.Printf("[ERROR] wrong number of variables for set command, need 2 (key value) but found %d", len(query)-1)
		return
	}

	db.Set(query[1], query[2])
}

func get(db *memdb.MemDB, query []string) {
	if len(query) != 2 {
		log.Printf("[ERROR] wrong number of variables for get command, need 1 (key) but found %d", len(query)-1)
		return
	}

	value := db.Get(query[1])

	fmt.Println(value)
}

func deleteEntry(db *memdb.MemDB, query []string) {
	if len(query) != 2 {
		log.Printf("[ERROR] wrong number of variables for delete command, need 1 (key) but found %d", len(query)-1)
		return
	}

	db.Delete(query[1])
}

func count(db *memdb.MemDB, query []string) {
	if len(query) != 2 {
		log.Printf("[ERROR] wrong number of variables for count command, need 1 (key) but found %d", len(query)-1)
		return
	}

	fmt.Println(db.Count(query[1]))
}

func readStdIn(reader *bufio.Reader) []string {
	fmt.Print("go-memdb=> ")
	input, _ := reader.ReadString('\n')

	// convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)
	return strings.Split(input, " ")
}
