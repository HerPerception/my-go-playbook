package main

type Employee struct {
	Name       string
	Department string
}

func GroupEmployees(EmployeeData []Employee) map[string][]string {
	EmployeesRegistry := make(map[string][]string)

	for _, value := range EmployeeData {
		var NamesSlice []string
		if employeeNames, exists := EmployeesRegistry[value.Department]; exists {
			employeeNames = append(employeeNames, value.Name)
			EmployeesRegistry[value.Department] = employeeNames
		} else {
			NamesSlice = append(NamesSlice, value.Name)
			EmployeesRegistry[value.Department] = NamesSlice
		}
	}
	return EmployeesRegistry
}

// func main() {
// 	input := []Employee{
// 		{Name: "Alice", Department: "HR"},
// 		{Name: "Bob", Department: "IT"},
// 		{Name: "Charlie", Department: "IT"},
// 		{Name: "David", Department: "HR"},
// 	}
// 	fmt.Println(GroupEmployees(input))
// }
