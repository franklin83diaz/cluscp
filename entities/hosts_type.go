package entities

type HostsType int

const (
	unknown HostsType = iota
	ip
	hostname
)

func (h HostsType) Set(s string) {
	if s == "I" || s == "i" {
		h = ip
	}
	if s == "H" || s == "h" {
		h = hostname
	}
	h = unknown
}
