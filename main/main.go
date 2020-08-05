package main

import (
"github.com/rexue2019/micro/cfg"
 )

type Config struct {
	cfg.ConfigurationImp
}

func main() {
	server.Run(&Config{}, func(microCtx microkit.MicroContext) {


	})
}