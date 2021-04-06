package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Summary struct {
	Status int    `json:"status"`
	Result Result `json:"result"`
}

type Result struct {
	Pst string  `json:"postcode"`
	Est int     `json:"eastings"`
	Nrt int     `json:"northings"`
	Lng float64 `json:"longitude"`
	Lat float64 `json:"latitude"`
}

func main() {
	// create array, 3 string elements
	postcodeArray := [3]string{"KT12DN", "EC3N4AB", "LE27FL"}
	// loop
	for i := 0; i < len(postcodeArray); i++ {
		res, err := http.Get("http://api.postcodes.io/postcodes/" + postcodeArray[i])
		if err != nil {
			panic(err.Error())
		}

		resData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err.Error())
		}

		var summary Summary
		// parse json-encoded data
		json.Unmarshal([]byte(resData), &summary)
		// write to file
		file, _ := json.MarshalIndent(summary.Result, "", "  ")
		ioutil.WriteFile(string(postcodeArray[i])+".json", file, 0644)

		fmt.Println(summary.Result)
	}
}
