package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'mostActive' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY customers as parameter.
 */

func mostActive(customers []string) []string {
	// Write your code here
	customerSlice := []string{}
	customerMap := make(map[string]int)
	for _, each := range customers {
		if val, exists := customerMap[each]; exists {
			customerMap[each] = val + 1
			// fmt.Println(each, val)
		} else {
			count := 1
			customerMap[each] = count
		}
	}
	// fmt.Println(customerMap)
	// for key, val := range customerMap {
	for r := 'A'; r <= 'Z'; r++ {
		for key, val := range customerMap {
			if rune(key[0]) == r {
				//fmt.Println(key)
				num := val
				//fmt.Println(num)
				p := float64((num * 100) / len(customers))
				//fmt.Println(p)
				if p < 5.0 {
					continue
				} else {
					for alp := 'a'; alp <= 'z'; alp++ {
						if rune(key[1]) == alp {
							customerSlice = append(customerSlice, key)
						}
					}
				}
			}
		}
	}

	for i := 0; i < len(customerSlice); i++ {
		if i < len(customerSlice)-1 && (customerSlice[i][0] == customerSlice[i+1][0]) && customerSlice[i][1] < customerSlice[i+1][1] {
			temp := customerSlice[i]
			customerSlice[i] = customerSlice[i+1]
			customerSlice[i+1] = temp
		}
	}
	return customerSlice
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	customersCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var customers []string

	for i := 0; i < int(customersCount); i++ {
		customersItem := readLine(reader)
		customers = append(customers, customersItem)
	}

	result := mostActive(customers)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
