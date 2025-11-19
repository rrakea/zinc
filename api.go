package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

const AUR_URL = "https://aur.archlinux.org/rpc/v5/"

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
	Out_of_date bool    `json:"OutOfDate"`
	Maintainer  string  `json:"Maintainer"`
}

func Search(name string) []Package {
	url := AUR_URL + "/search/" + name + "?by=name"
	res := send_request(url)
	if res.Res_type != "search" {
		log.Fatal("Incorrect API response type")
	}

	for _, item := range res.Body {
		item.print()
		fmt.Println()
		fmt.Println()
	}

	return res.Body
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
		log.Fatal("Not ok http status")
	}

	if res.Version != 5 {
		log.Fatal("Version other than 5 not supported. Got version: ", res.Version)
	}

	if res.Res_type == "error" {
		log.Fatal("API Response Error: ", res.Error)
	}

	return res
}

func Info(name string) {

}

func (r *Package) print() {
	fmt.Println("Name: " + r.Name)
	fmt.Println("Version: " + r.Version)
	fmt.Println("Desc: " + r.Desc)
	fmt.Println("URL: " + r.Url)
	fmt.Println("Votes: " + strconv.Itoa(r.Votes))
	fmt.Println("Popularity: " + strconv.FormatFloat(r.Popularity, 'e', -1, 64))
	fmt.Println("Out of date: " + strconv.FormatBool(r.Out_of_date))
	fmt.Println("Maintainer: " + r.Maintainer)
}
