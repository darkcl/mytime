package server

import "github.com/darkcl/mytime/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.port"))
}
