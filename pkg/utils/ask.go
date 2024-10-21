package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func AskOptions(s string, want []string) string {
	fmt.Println(s)
	var resp string
	fmt.Scanln(&resp)

	resp = strings.TrimSpace(resp)

	if !Contains(want, resp) {
		fmt.Println("Invalid option")
		return AskOptions(s, want)
	}
	return resp
}

func AskMinLength(s string, minLength int) string {
	fmt.Println(s)
	var resp string
	fmt.Scanln(&resp)

	resp = strings.TrimSpace(resp)

	if len(resp) < minLength {
		if len(resp) == 0 {
			fmt.Println("Your answer is empty.")
		} else {
			fmt.Println("Minimum length is", minLength)
		}
		return AskMinLength(s, minLength)
	}
	return resp
}

func AskNumber(s string, min int, max int) int {
	fmt.Println(s)
	var resp string
	fmt.Scanln(&resp)

	resp = strings.TrimSpace(resp)

	number, err := strconv.Atoi(resp)
	if err != nil {
		fmt.Println("Invalid number")
		return AskNumber(s, min, max)
	}

	if number < min || number > max {
		fmt.Println("Number out of range")
		return AskNumber(s, min, max)
	}

	return number
}
