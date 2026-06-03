package main

func LRUCache(newPage string) []string {
	var pages []string
	if len(pages) >= 3 {
		pages = append(pages, newPage)
		return pages[:len(pages)-3]
	} else if len(pages) < 3 {
		pages = append(pages, newPage)
	}
	return pages
}

// func main() {

// }
