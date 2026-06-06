package main

import "fmt"

func DescendComb() {
	for i := 9; i >= 0; i-- {
		for j := 9; j >= i; j-- {
			for k := 9; k >= 0; k-- {
				for l := j - 1; l >= 0; l-- {
					if i == 0 && j == 1 && k == 0 && l == 0 {
						fmt.Printf("%d%d %d%d", i, j, k, l)
						break
					}
					fmt.Printf("%d%d %d%d, ", i, j, k, l)
				}
			}
		}
	}
}

func main() {
	DescendComb()
}
