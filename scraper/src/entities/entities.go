package entities

type Record struct {
	Port                    string `json:"port,omitempty"`
	Status                  string `json:"status,omitempty"`
	SentPacketsLifetime     int64  `json:"sent_packets_lifetime,omitempty"`
	ReceivedPacketsLifetime int64  `json:"received_packets_lifetime,omitempty"`
	Collisions              int64  `json:"collisions,omitempty"`
	SentBytesPerSecond      int64  `json:"sent_bytes_per_second,omitempty"`
	ReceivedBytesPerSecond  int64  `json:"received_bytes_per_second,omitempty"`
	UpTime                  string `json:"up-time,omitempty"`
}
