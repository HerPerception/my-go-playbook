package main

func ClassifyAge(age int) string {
	switch {
	case age < 18:
		return "minor"

	case age >= 18 && age <= 64:
		return "adult"

	case age >= 65:
		return "senior"
	}
	return ""
}

// func main() {
// 	fmt.Println(ClassifyAge(45))
// }
