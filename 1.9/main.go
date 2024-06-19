package main

import (
	"fmt"
)

func main() {
	var year int
	_, _ = fmt.Scan(&year)
	fmt.Print(checkYear(year))
}

func checkYear(year int) string {
	if year%400 == 0 || year%4 == 0 && year%100 != 0 {
		return "YES"
	} else {
		return "NO"
	}
}
