package main

import (
	"context"
	"fmt"
	"github.com/rexue2019/k8sdemo/rpc"
	"github.com/rexue2019/micro/cfg"
 	"github.com/rexue2019/micro/server"
)

type Config struct {
	cfg.ConfigurationImp
}

type demoImp struct {
}

func (this *demoImp) HelloNoReponse(_ context.Context, req  *rpc.HelloRequest) error {
	fmt.Println("receive", req)
	return nil
}

func main() {
	server.Run(&Config{}, func(microCtx server.MicroContext) {
		rpc.RegisterK8sdemoForHTTP(&demoImp{})
	})
}
