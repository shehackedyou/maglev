package maglev

import (
	"errors"

	server "github.com/multiverse-os/maglev/server"
)

func (f *Framework) HTTP() server.Server {
	var err error
	if f.servers[server.HTTP] == nil {
		f.servers[server.HTTP], err = server.New(server.HTTP, f.Config.Address, f.Config.Port)
		if err != nil {
			panic(errors.New("failed to create new http server:" + err.Error()))
		}
	}
	return f.servers[server.HTTP]
}
