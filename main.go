package main

import (
	"fmt"
	"strings"
)

const original = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz 123456789"

func hashed(key int, str string) string {
	runes := []rune(str)
	lastletter := string(runes[len(str)-key : len(str)])
	left := string(runes[:len(str)-key])
	return fmt.Sprintf("%s%s", lastletter, left)

}

func encrpt(key int, str string) string {
	hashletter := hashed(key, original)
	var hashedstr = ""
	findone := func(r rune) rune {
		pos := strings.Index(original, string([]rune{r}))
		if pos != -1 {
			hashedstr = hashedstr + string(hashletter[pos])
			return r
		}
		return r

	}
	strings.Map(findone, str)
	return hashedstr

}
func decrypt(key int, str string) string {
	hashletter := hashed(key, original)
	var hashedstr = ""
	findone := func(r rune) rune {
		pos := strings.Index(hashletter, string([]rune{r}))
		if pos != -1 {
			hashedstr = hashedstr + string(original[pos])
			return r
		}
		return r
	}

	strings.Map(findone, str)
	return hashedstr

}

func main() {
	var str string
	fmt.Scanf("%s", &str)
	key := 5
	D := encrpt(key, str)
	fmt.Println(D)
	fmt.Println(decrypt(key, D))

}
