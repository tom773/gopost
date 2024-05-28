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
	Endp   string
}

func parseArgs() Args {

	var args Args

	JSONflag := flag.String("j", "", "data")
	Methodflag := flag.String("m", "", "method")
	Endpflag := flag.String("e", "", "endpoint")

	flag.Parse()

	args.JSON = *JSONflag
	args.Method = *Methodflag
	args.Endp = *Endpflag
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
	parsed := parseJson(args)

	sendRequest(args, parsed)
}

func parseJson(args Args) []byte {

	split := strings.Split(args.JSON, ":")
	jsonbody := make(map[string]string)
	jsonbody[split[0]] = split[1]
	jsonData, err := json.Marshal(jsonbody)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(colors.Green, "Body: ", colors.Yellow, string(jsonData), colors.Reset)
	return jsonData
}

// This will be a flag at some point
type Request struct {
	UserID string `json:"userid"`
}

func sendRequest(args Args, jsonData []byte) {
	client := resty.New()
	err := json.Unmarshal(jsonData, &Request{})
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsonData).
		Execute(args.Method, "http://"+args.Url+":"+args.Port+"/"+args.Endp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(colors.Green, "Response: ", colors.Yellow, resp, colors.Reset)
}
