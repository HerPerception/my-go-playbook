package main

func Swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// func main() {
// 	a := 10
// 	b := 5
// 	Swap(&a, &b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

/* In this exercise, in the function signature, *int is a type - pointer to an int,
not an expression. It points to the memory address of an integer.
In *n, an axpression, * means dereference - follow the address stored in n and access the value there,
so *n += 1 is go to the address of n and increment the value by 1.
Dereferencing is the programiing process of using a pointer to access or modify the value stored at the specific memory address it points to.*/
