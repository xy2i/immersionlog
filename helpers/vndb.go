package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://api.vndb.org/kana/vn"

type VN struct {
	More    bool `json:"more"`
	Results []struct {
		//		Alttitle string `json:"alttitle"`
		ID    string `json:"id"`
		Title string `json:"title"`
	} `json:"results"`
}

// Work in progress, this function should return a []discordgo.ApplicationCommandOptionChoice, in its current state for testing.
func searchVN(query string) {
	data := `{"filters": ["search", "=", "` + query + `"], "fields": "title", "sort": "searchrank"}`
	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var result VN
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	for _, rec := range result.Results {
		fmt.Println(rec.Title)
	}
}
