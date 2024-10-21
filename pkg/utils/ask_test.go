package utils

import (
	"testing"
	"time"
)

func TestAskOptions(t *testing.T) {
	type args struct {
		s    string
		want []string
	}
	tests := []struct {
		name   string
		args   args
		want   string
		inputs []string
	}{
		{"AskOptions", args{"Do you want to use IP or Hostname? (I/H)", []string{"I", "H", "i", "h"}}, "I", []string{"m", "I"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mocking the stdin
			StdinMocking(tt.inputs, 500*time.Millisecond)

			// Test the function
			if got := AskOptions(tt.args.s, tt.args.want); got != tt.want {
				t.Errorf("AskOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAskMinLength(t *testing.T) {
	type args struct {
		s         string
		minLength int
	}
	tests := []struct {
		name   string
		args   args
		want   string
		inputs []string
	}{
		{"AskMinLength", args{"Enter the file name to generate:", 1}, "file", []string{"", "file"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mocking the stdin
			StdinMocking(tt.inputs, 500*time.Millisecond)
			if got := AskMinLength(tt.args.s, tt.args.minLength); got != tt.want {
				t.Errorf("AskMinLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAskNumber(t *testing.T) {
	type args struct {
		s   string
		min int
		max int
	}
	tests := []struct {
		name   string
		args   args
		want   int
		inputs []string
	}{
		{"AskNumber", args{"Enter the number:", 1, 10}, 5, []string{"5"}},
		{"AskNumber", args{"Enter the number:", 1, 10}, 5, []string{"", "5"}},
		{"AskNumber", args{"Enter the number:", 1, 10}, 5, []string{"no number", "5"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mocking the stdin
			StdinMocking(tt.inputs, 500*time.Millisecond)
			if got := AskNumber(tt.args.s, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("AskNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
