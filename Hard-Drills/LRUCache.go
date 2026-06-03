package main

type LRUCache struct {
	History []string
}

// func LRUCache(newPage string) []string {
// 	var pages []string
// 	if len(pages) >= 3 {
// 		pages = append(pages, newPage)
// 		return pages[:len(pages)-3]
// 	} else if len(pages) < 3 {
// 		pages = append(pages, newPage)
// 	}
// 	return pages
// }

func (c *LRUCache) Visit(newPage string) []string {
	c.History = append(c.History, newPage)
	//var history []string
	if len(c.History) > 3 {
		//history = append(history, c.History[len(c.History)-3:]...)
		c.History = c.History[len(c.History)-3:]
	}
	return c.History
}

// func main() {
// 	cache := &LRUCache{}
// 	cache.Visit("home")
// 	cache.Visit("about")
// 	cache.Visit("contact")
// 	cache.Visit("pricing")
// 	fmt.Println(cache.History)
// 	fmt.Println(cache.Visit)
// }
