package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	today := currentTime.Format("02-01-2006")
	// fmt.Println(today - 1)
	fmt.Println(today[:2])
	t := time.Now().AddDate(0, 0, -1).Format("02-01-2006")
	fmt.Println(t)
	// colorReset := "\033[0m"

	// colorRed := "\033[31m"
	// colorGreen := "\033[32m"
	// colorYellow := "\033[33m"
	// colorBlue := "\033[34m"
	// colorPurple := "\033[35m"
	// colorCyan := "\033[36m"
	// colorWhite := "\033[37m"

	// fmt.Printf("%s%s%d\n", string(colorRed), " test ", 4)
	// fmt.Println(string(colorGreen), "test")
	// fmt.Println(string(colorYellow), "test")
	// fmt.Println(string(colorBlue), "test")
	// fmt.Println(string(colorPurple), "test")
	// fmt.Println(string(colorWhite), "test")
	// fmt.Println(string(colorCyan), "test", string(colorReset))
	// fmt.Println("next")
}
