package bungee

type BungeePlayer struct {
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	IP     string `json:"-"`
	Server string `json:"server"`
}
