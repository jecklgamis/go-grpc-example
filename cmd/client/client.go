package main

import (
	"flag"
	"fmt"
	"github.com/jecklgamis/grpc-go-example/pkg/client"
	"log"
	"os"
)

var (
	serverAddr = flag.String("serverAddr", "localhost:4000", "The server address")
)

func printUsageAndExit() {
	flag.Usage()
	os.Exit(-1)
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		printUsageAndExit()
	}
	cmd := flag.Arg(0)
	switch cmd {
	case "get":
		kvClient := client.New(*serverAddr)
		if len(flag.Args()) != 2 {
			println("Get requires key")
			os.Exit(-1)
		}
		k := flag.Args()[1]
		v, err := kvClient.Get(k)
		if err != nil {
			log.Fatal(err)
		}
		if v != "" {
			println(v)
		} else {
			println("No value found for key", k)
		}
	case "put":
		kvClient := client.New(*serverAddr)
		if len(flag.Args()) != 3 {
			println("Put requires key value")
			os.Exit(-1)
		}
		k := flag.Args()[1]
		v := flag.Args()[2]
		err := kvClient.Put(k, v)
		if err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("Unknown command", cmd)
	}
}
