package main

import "fmt"

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	foodMap := make(map[string]int)
	var time food
	foodMap = map[string]int{
		"burger":  15,
		"chips":   10,
		"nuggets": 12,
	}
	if cookTime, exists := foodMap[order]; exists {
		time.preptime = cookTime
	} else {
		return 404
	}
	return time.preptime
}
func main() {
	fmt.Println(FoodDeliveryTime("burger"))
	fmt.Println(FoodDeliveryTime("chips"))
	fmt.Println(FoodDeliveryTime("nuggets"))
	fmt.Println(FoodDeliveryTime("burger") + FoodDeliveryTime("chips") + FoodDeliveryTime("nuggets"))
}
