package pkg

import (
	"os"
	"reflect"
	"testing"
)

func TestReadHosts(t *testing.T) {
	createHostsFile("hosts_test", []string{"192.0.2.1", "192.0.2.2"})
	type args struct {
		fileHost string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"Read hosts from file", args{"hosts_test"}, []string{"192.0.2.1", "192.0.2.2"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadHosts(tt.args.fileHost)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadHosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadHosts() = %v, want %v", got, tt.want)
			}
		})
	}
	// Remove the file
	err := os.Remove("hosts_test")
	if err != nil {
		t.Error(err)
	}
}
