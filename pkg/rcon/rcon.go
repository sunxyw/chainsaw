package rcon

type client struct {
	servers map[string]*server
}

var Client *client

func InitRconClient() {
	Client = &client{
		servers: make(map[string]*server),
	}
}

func AddServer(conf ServerConf) {
	Client.servers[conf.Name] = newServer(conf)
	Client.servers[conf.Name].Connect()
}

func Server(name string) *server {
	return Client.servers[name]
}
