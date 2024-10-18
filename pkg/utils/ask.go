package utils

import (
	"fmt"
	"strings"
)

func AskOptions(s string, want []string) string {
	fmt.Println(s)
	var resp string
	fmt.Scanln(&resp)

	resp = strings.TrimSpace(resp)

	if !Contains(want, resp) {
		AskOptions(s, want)
		fmt.Println("Invalid option")
		return resp
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
		AskMinLength(s, minLength)
		return resp
	}
	return resp
}
