package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	h := n / 30
	m := (n % 30) * 2
	fmt.Println("It is", h, "hours", m, "minutes.")
}
