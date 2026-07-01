package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func main() {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	var artists []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, each := range artists {
		fmt.Printf("Name:%s, Creation Date:%d\n\n", each.Name, each.CreationDate)
	}

	fmt.Println(len(artists))
}
