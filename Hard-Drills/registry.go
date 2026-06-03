package main

// func CheckUser(registryMap map[string]bool, key string) string {
// 	if active, ok := registryMap[key]; !ok || !active {
// 		return "Invalid or inactive username"
// 		//registryMap[key] = true
// 	}
// 	//if active, _ := registryMap[key]; active {
// 	return "User active"
// 	//	}
// 	//
// 	// fmt.Println(ok)
// }

func CheckUser(registryMap map[string]bool, key string) (bool, bool) {
	active, ok := registryMap[key]
	return active, ok
}

// func main() {
// 	registry := map[string]bool{
// 		"oofforbu": true,
// 		"vintech":  true,
// 		"codo":     true,
// 		"jadah":    true,
// 	}

// 	Registry(CheckUser, "victor")
// }
