package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/go-resty/resty/v2"
)

type Colors struct {
	Reset  string
	Red    string
	Green  string
	Yellow string
	Blue   string
	Purple string
	Cyan   string
	White  string
}

var colors = Colors{
	Reset:  "\033[0m",
	Red:    "\033[31m",
	Green:  "\033[32m",
	Yellow: "\033[33m",
	Blue:   "\033[34m",
	Purple: "\033[35m",
	Cyan:   "\033[36m",
	White:  "\033[37m",
}

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

	args.JSON = *JSONflag
	args.Method = *Methodflag
	if len(flag.Args()) < 2 {
		log.Fatal("Usage: program [flags] <url> <port>")
	}
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
	fmt.Println(colors.Green, "\nServer is healthy!\n", colors.Reset)

	args := parseArgs()
	fmt.Println(colors.Purple, args.Method, colors.Green, "\n", args.Url, ":", colors.Blue, args.Port, colors.Reset)
	parseJson(args)
}

type JsonBody struct {
	Key   string
	Value string
}

func parseJson(args Args) {

	split := strings.Split(args.JSON, ":")
	jsonbody := make(map[string]string)
	jsonbody[split[0]] = split[1]
	jsonData, err := json.Marshal(jsonbody)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(colors.Cyan, "Body: ", colors.Yellow, string(jsonData), colors.Reset)
}
