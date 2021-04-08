package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/antchfx/xmlquery"
)

var (
	file = flag.String("file", "", "path to xml file")
)

func abortOnErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

type Ids struct {
	Values []string
}

func (ids *Ids) Append(id string) bool {
	found := false
	for _, v := range ids.Values {
		if v == id {
			found = true
			break
		}
	}
	if !found {
		ids.Values = append(ids.Values, id)
	}
	return !found
}

func main() {

	flag.Parse()
	var ids Ids

	f, err := os.Open(*file)
	abortOnErr(err)

	start := time.Now()

	p, err := xmlquery.CreateStreamParser(f,
		"/PublicationDelivery/dataObjects/SiteFrame/stopPlaces/StopPlace")

	abortOnErr(err)
	for {
		node, err := p.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			abortOnErr(err)
		}
		for _, a := range node.Attr {
			if a.Name.Local == "id" {
				appended := ids.Append(a.Value)
				if !appended {
					fmt.Printf("%s[%s=%s]\n",
						node.Data, a.Name.Local, a.Value)
				}
			}
		}
	}
	duration := time.Since(start)
	fmt.Println("Duration:", duration)
}
