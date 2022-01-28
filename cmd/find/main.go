package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tengla/nsr-parser/stopplace"
)

var (
	file       = flag.String("file", "./nsr.current.json", "json source file")
	parameters = flag.String("parameters", "", "json-struct containing search parameters")
	jsonout    = flag.Bool("jsonout", false, "JSON output")
)

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getParams() stopplace.ParamMap {
	var params stopplace.SearchParams
	json.Unmarshal([]byte(*parameters), &params)
	return params.ExtractParams()
}

func getScanner() *bufio.Scanner {
	f, err := os.Open(*file)
	failOnError(err)
	scanner := bufio.NewScanner(f)
	// detect newline
	onNewLine := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == '\n' /* 10 newline*/ {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onNewLine)
	return scanner
}

// color table
var colors = map[string]string{
	"reset":  "\033[0m",
	"red":    "\033[31m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"blue":   "\033[34m",
	"purple": "\033[35m",
	"cyan":   "\033[36m",
	"white":  "\033[37m",
}

// printStop - print StopPlace to stdout
func printStop(sp stopplace.StopPlace) {
	if *jsonout {
		bytes, err := json.Marshal(sp)
		failOnError(err)
		fmt.Println(string(bytes))
		fmt.Println("")
		return
	}
	fmt.Println("")
	fmt.Printf("%sStopPlace(%sID='%s', Name='%s', ", colors["white"], colors["purple"], sp.ID, sp.Name.Text)
	fmt.Printf("Lon=%s, Lat=%s, %skeylist=[",
		sp.Centroid.Location.Longitude, sp.Centroid.Location.Latitude, colors["yellow"])
	keyvals := make([]string, 0)
	for _, kv := range sp.KeyList.KeyValue {
		item := fmt.Sprintf("%s='%s'", kv.Key, kv.Value)
		keyvals = append(keyvals, item)
	}
	fmt.Printf("%s]%s)%s\n",
		strings.Join(keyvals, ", "),
		colors["white"],
		colors["reset"])
}

func main() {
	flag.Parse()
	params := getParams()
	scanner := getScanner()
	for scanner.Scan() {
		var sp stopplace.StopPlace
		bytes := scanner.Bytes()
		if len(bytes) == 0 /*we hit the end*/ {
			continue
		}
		err := json.Unmarshal(bytes, &sp)
		if err != nil {
			fmt.Printf("error parsing json %s\n", err.Error())
			continue
		}
		if stopplace.Match(params, sp) {
			printStop(sp)
		}
	}
}
