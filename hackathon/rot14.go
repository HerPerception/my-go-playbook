package main

func Rot14(s string) string {
	var newstr string
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			alp := ch - 'A'
			rot := (alp + 14)
			if rot >= 26 {
				rot = rune(rot - 26)
			}
			newstr += string(rune(rot + 'A'))
		} else if ch >= 'a' && ch <= 'z' {
			alp := ch - 'a'
			rot := rune(alp + 14)
			if rot >= 26 {
				rot = rune(rot - 26)
			}
			newstr += string(rune(rot + 'a'))
		} else {
			newstr += string(ch)
		}
	}
	return newstr
}

// func main() {
// 	result := Rot14("Hello! How are You?")

// 	for _, r := range result {
// 		z01.PrintRune(r)
// 	}
// 	z01.PrintRune('\n')
// }
