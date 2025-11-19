package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const AUR_URL = "https://aur.archlinux.org/rpc/v5/"

var search_chan chan []Package

type API_return struct {
	Version      int       `json:"version"`
	Res_type     string    `json:"type"`
	Result_count int       `json:"resultcount"`
	Body         []Package `json:"results"`
	Error        string    `json:"error"`
}

type Package struct {
	Name        string  `json:"Name"`
	Version     string  `json:"Version"`
	Desc        string  `json:"Description"`
	Url         string  `json:"URL"`
	Votes       int     `json:"NumVotes"`
	Popularity  float64 `json:"Popularity"`
	Out_of_date int     `json:"OutOfDate"`
	Maintainer  string  `json:"Maintainer"`
}

func Search(name string) {
	url := AUR_URL + "/search/" + name + "?by=name"
	res := send_request(url)
	if res.Res_type != "search" {
		log.Fatal("Incorrect API response type")
	}
	search_chan <- res.Body
}

func send_request(url string) API_return {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Could access the aur; ", err)
	}
	defer resp.Body.Close()

	var res API_return
	switch resp.StatusCode {
	case http.StatusOK:
		{
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			err = json.Unmarshal(body, &res)
			if err != nil {
				log.Fatal(err)
			}
		}
	default:
		fmt.Println("HTTP Code: ", resp.StatusCode)
	}

	if res.Version != 5 {
		log.Fatal("Version other than 5 not supported. Got version: ", res.Version)
	}

	if res.Res_type == "error" {
		fmt.Println("API Response Error: ", res.Error)
	}

	return res
}
