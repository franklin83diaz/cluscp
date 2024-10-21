package utils

import (
	"os"
	"time"
)

func StdinMocking(inputs []string, Millisecond time.Duration) {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdin = r
	go func() {
		defer w.Close()
		for _, input := range inputs {
			time.Sleep(500 * time.Millisecond)
			w.Write([]byte(input + "\n"))
		}
	}()
}
