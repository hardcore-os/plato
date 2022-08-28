package discov

type Service struct {
	Name      string      `json:"name"`
	Endpoints []*Endpoint `json:"endpoints"`
}

type Endpoint struct {
	ServerName string `json:"server_name"`
	IP         string `json:"ip"`
	Port       int    `json:"port"`
	Weight     int    `json:"weight"`
	Enable     bool   `json:"enable"`
}
