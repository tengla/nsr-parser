package stopplace

import (
	"log"
	"reflect"
	"regexp"
	"strconv"

	geo "github.com/kellydunn/golang-geo"
)

type SearchParams struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	TransportMode string  `json:"transportmode"`
	Key           string  `json:"key"`
	Value         string  `json:"value"`
	Longitude     float64 `json:"longitude"`
	Latitude      float64 `json:"latitude"`
	Distance      float64 `json:"distance"`
}

type Location struct {
	Longitude float64
	Latitude  float64
	Distance  float64
}

// MatchLocation
func (location *Location) MatchLocation(sp StopPlace) bool {
	point := geo.NewPoint(location.Latitude, location.Longitude)
	lat, err := strconv.ParseFloat(sp.Centroid.Location.Latitude, 32)
	if err != nil {
		log.Fatal(err)
		return false
	}
	lon, err := strconv.ParseFloat(sp.Centroid.Location.Longitude, 32)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return (point.GreatCircleDistance(geo.NewPoint(lat, lon)) <= location.Distance)
}

type ParamMap map[string]interface{}

// table of truth
var table = map[string]func(params ParamMap, sp StopPlace) bool{
	"ID": func(params ParamMap, sp StopPlace) bool {
		id, exists := params["ID"]
		if !exists {
			return false
		}
		return id.(string) == sp.ID
	},
	"Key": func(params ParamMap, sp StopPlace) bool {
		for _, kv := range sp.KeyList.KeyValue {
			if kv.Key == params["Key"] && kv.Value == params["Value"] {
				return true
			}
		}
		return false
	},
	"Value": func(params ParamMap, sp StopPlace) bool {
		return true
	},
	"Name": func(params ParamMap, sp StopPlace) bool {
		value, exists := params["Name"]
		if !exists {
			return false
		}
		name := value.(string)
		re, err := regexp.Compile("(?i)" + name)
		if err != nil {
			log.Fatal(err)
			return false
		}
		return re.MatchString(sp.Name.Text)
	},
	"Longitude": func(params ParamMap, sp StopPlace) bool {
		l := &Location{}
		for k, v := range params {
			if k == "Longitude" {
				l.Longitude = v.(float64)
			}
			if k == "Latitude" {
				l.Latitude = v.(float64)
			}
			if k == "Distance" {
				l.Distance = v.(float64)
			}
		}
		return l.MatchLocation(sp)
	},
	"Latitude": func(params ParamMap, sp StopPlace) bool {
		return true
	},
	"Distance": func(params ParamMap, sp StopPlace) bool {
		return true
	},
	"TransportMode": func(params ParamMap, sp StopPlace) bool {
		v, exists := params["TransportMode"]
		if !exists {
			return false
		}
		return v == sp.TransportMode
	},
}

func (s *SearchParams) ExtractParams() map[string]interface{} {
	v := reflect.ValueOf(*s)
	t := reflect.TypeOf(*s)
	collected := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch field.Type().Name() {
		case "float64":
			value := field.Float()
			if value != 0 {
				collected[t.Field(i).Name] = value
			}
		case "string":
			value := field.String()
			if len(value) > 0 {
				collected[t.Field(i).Name] = value
			}
		}
	}
	return collected
}

// Match
func Match(params ParamMap, sp StopPlace) bool {
	count := 0
	for k := range params {
		truthy, exists := table[k]
		if exists && truthy(params, sp) {
			count += 1
		}
	}
	return count >= len(params)
}
