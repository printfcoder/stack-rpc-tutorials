package config

var (
	Topic = "mu.micro.nsq.demo"

	NsqLookupdAddrs = []string {"172.16.58.20:4161"}
	NsqdAddrs = []string {"172.16.58.20:4150"}
	NsqMaxInFlight = 5
)
