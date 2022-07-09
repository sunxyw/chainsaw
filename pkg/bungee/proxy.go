package bungee

type BungeeProxy struct {
	Name string
}

func NewBungeeProxy(name string) *BungeeProxy {
	return &BungeeProxy{
		Name: name,
	}
}
