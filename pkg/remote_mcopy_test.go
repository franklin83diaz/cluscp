package pkg_test

import (
	"cluscp/pkg"
	"reflect"
	"testing"

	"golang.org/x/crypto/ssh"
)

func TestConnectSSH(t *testing.T) {
	type args struct {
		listHost []string
		authMod  string
	}
	tests := []struct {
		name         string
		args         args
		wantSessions []*ssh.Session
		wantErr      bool
	}{
		{"Test1", args{[]string{"127.0.0.1"}, "password"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSessions, err := pkg.ConnectSSH(tt.args.listHost, tt.args.authMod)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectSSH() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSessions, tt.wantSessions) {
				t.Errorf("ConnectSSH() = %v, want %v", gotSessions, tt.wantSessions)
			}
		})
	}
}
