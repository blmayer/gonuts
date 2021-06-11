package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	nats "github.com/nats-io/nats.go"
)

const help = `gonuts makes a request to your nats server
Usage:
  gonuts [options]
Available options:
  -h
  --help	show this help
Examples:
  gonuts --help`


func main() {
	addr := os.Args[1]
	subj := os.Args[2]
	
	// Command line arguments parsing
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-h":
			fallthrough
		case "--help":
			println(help)
			os.Exit(0)
		}
	}
	
	// Connect to nats
	conn, err := nats.Connect(addr)
	if err != nil {
		log.Fatalln(err)
	}

	timeout := time.Duration(60 * time.Second)
	data := []byte(os.Args[3])
	if os.Args[3] == "-f" {
		file, err := os.Open(os.Args[4])
		data, err = ioutil.ReadAll(file)
		if err != nil {
			log.Fatalln("file read error:", err)
		}
	}

	res, err := conn.Request(subj, data, timeout)
	if err != nil {
		log.Fatalln("request error:", err)
	}

	println(string(res.Data))
	conn.Close()
}
