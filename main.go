package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
)

type Args struct {
	Url    string
	Port   string
	Method string
	JSON   string
}

func parseArgs() Args {

	var args Args

	JSONflag := flag.String("j", "", "data")
	Methodflag := flag.String("m", "", "method")

	flag.Parse()
	fmt.Println(*JSONflag)
	fmt.Println(*Methodflag)

	args.JSON = *JSONflag
	args.Method = *Methodflag
	args.Url = flag.Arg(0)
	args.Port = flag.Arg(1)

	return args
}

func main() {
	client := resty.New()
	_, err := client.R().Get("http://localhost:8090/api/health")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is healthy!")

	args := parseArgs()
	fmt.Fprintf(os.Stdout, "%s\nTesting: %s on Port: %s\nBody: %s", args.Method, args.Url, args.Port, args.JSON)
}
