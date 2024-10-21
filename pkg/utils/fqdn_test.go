package utils

import (
	"reflect"
	"testing"
)

func TestGetFqdnList(t *testing.T) {
	type args struct {
		hostname string
		start    int
		end      int
		domain   string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"Test 1", args{"test", 1, 3, "com"}, []string{"test1.com", "test2.com", "test3.com"}, false},
		{"Test 2", args{"test", 3, 1, "com"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFqdnList(tt.args.hostname, tt.args.start, tt.args.end, tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFqdnList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFqdnList() = %v, want %v", got, tt.want)
			}
		})
	}
}
