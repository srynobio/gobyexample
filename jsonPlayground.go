package main

import (
	"encoding/json"
	"fmt"
)

type Cars struct {
	Make, Model string
	Year        int
	EngineSize  []int
	Version     map[string]string
}

var pt = fmt.Println

func main() {

	block := []int{225, 318, 360, 440}
	versions := map[string]string{
		"Dart":       "Demon",
		"Duster":     "Swinger",
		"Challenger": "Hemi",
	}

	c := Cars{"Dodge", "Dart", 1973, block, versions}
	b, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	pt(string(b))

	var ride Cars
	myRide := json.Unmarshal(b, &ride)
	if myRide != nil {
		panic(myRide)
	}
	pt(ride)
	pt(ride.Year)
	for v, k := range ride.Version {
		pt(k, ":----:", v)
	}
	pt(ride.Version["Duster"])
}
