package main

import "fmt"

type Artist struct {
	Name   string
	Date   string
	Member bool
}

func main() {
	Artists := []Artist{
		{"Ada Ehi", "12/02/2015", true},
		{"Mon Dee", "15/03/1990", false},
		{"Fred Stone", "08/10/2000", true},
	}

	for _, artist := range Artists {
		fmt.Println(artist.Name, artist.Date)
	}
	fmt.Println()
	Artists = append(Artists, Artist{"Jay Mich", "01/05/2004", false})
	for _, artist := range Artists {
		fmt.Println(artist.Name, artist.Date)
	}
}
