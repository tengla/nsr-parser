package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tengla/nsr-parser/stopplace"
)

var (
	key = flag.String("key", "", "The key to look for")
	val = flag.String("value", "", "The value to look for")
)

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printStop(sp stopplace.StopPlace) {
	fmt.Println("")
	fmt.Printf("ID='%s', Name='%s'\n ", sp.ID, sp.Name.Text)
	for _, kv := range sp.KeyList.KeyValue {
		fmt.Printf("%s='%s', ", kv.Key, kv.Value)
	}
	fmt.Println("")
}

func main() {
	flag.Parse()
	f, err := os.Open("./nsr.current.json")
	failOnError(err)
	scanner := bufio.NewScanner(f)
	onNewLine := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == '\n' {
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
		err := json.Unmarshal(scanner.Bytes(), &sp)
		if err != nil {
			continue
		}
		if len(*key) > 0 && len(*val) < 1 {
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
