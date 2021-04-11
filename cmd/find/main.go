package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/tengla/nsr-parser/stopplace"
)

var (
	file    = flag.String("file", "./nsr.current.json", "json source file")
	key     = flag.String("key", "", "The key to look for")
	val     = flag.String("value", "", "The value to look for")
	jsonout = flag.Bool("jsonout", false, "JSON output")
)

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printStop(sp stopplace.StopPlace) {
	if *jsonout {
		bytes, err := json.Marshal(sp)
		failOnError(err)
		fmt.Println(string(bytes))
		fmt.Println("")
		return
	}
	fmt.Println("")
	fmt.Printf("ID='%s', Name='%s'\n ", sp.ID, sp.Name.Text)
	for _, kv := range sp.KeyList.KeyValue {
		fmt.Printf("%s='%s', ", kv.Key, kv.Value)
	}
	fmt.Println("")
}

func main() {
	flag.Parse()
	f, err := os.Open(*file)
	failOnError(err)
	scanner := bufio.NewScanner(f)
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
		if *key == "ID" {
			printStop(sp)
		} else if *key == "Name" && len(*val) > 0 {
			re, err := regexp.Compile(*val)
			failOnError(err)
			if re.MatchString(sp.Name.Text) {
				printStop(sp)
			}
		} else if len(*key) > 0 && len(*val) < 1 {
			for _, kv := range sp.KeyList.KeyValue {
				if kv.Key == *key {
					printStop(sp)
				}
			}
		} else if len(*key) > 0 && len(*val) > 0 {
			for _, kv := range sp.KeyList.KeyValue {
				if kv.Key == *key && kv.Value == *val {
					printStop(sp)
				}
			}
		} else {
			printStop(sp)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
