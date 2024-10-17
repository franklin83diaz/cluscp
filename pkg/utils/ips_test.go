package utils

import (
	"math/big"
	"net"
	"testing"
)

func Test_ipToInt(t *testing.T) {
	type args struct {
		ip net.IP
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{"Test ipToInt IP 0.0.0.0", args{net.ParseIP("0.0.0.0")}, big.NewInt(0)},
		{"Test ipToInt IP 0.0.0.1", args{net.ParseIP("0.0.0.1")}, big.NewInt(1)},
		{"Test ipToInt IP 10.0.0.1", args{net.ParseIP("10.0.0.1")}, big.NewInt(167772161)},
		{"Test ipToInt IP 255.255.255.255", args{net.ParseIP("255.255.255.255")}, big.NewInt(4294967295)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ipToInt(tt.args.ip); got.Cmp(tt.want) != 0 {

				t.Errorf("ipToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intToIP(t *testing.T) {
	type args struct {
		ipInt *big.Int
	}
	tests := []struct {
		name string
		args args
		want net.IP
	}{
		{"Test intToIP IP 0.0.0.0", args{big.NewInt(0)}, net.ParseIP("0.0.0.0")},
		{"Test intToIP IP 0.0.0.1", args{big.NewInt(1)}, net.ParseIP("0.0.0.1")},
		{"Test intToIP IP 10.0.0.1", args{big.NewInt(167772161)}, net.ParseIP("10.0.0.1")},
		{"Test intToIP IP 255.255.255.255", args{big.NewInt(4294967295)}, net.ParseIP("255.255.255.255")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToIP(tt.args.ipInt); !got.Equal(tt.want) {
				t.Errorf("intToIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
