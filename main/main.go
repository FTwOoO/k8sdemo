package main

import (
	"context"
	"fmt"
	"github.com/FTwOoO/k8sdemo/rpc"
	"github.com/FTwOoO/micro/cfg"
	"github.com/FTwOoO/micro/server"
)

type Config struct {
	cfg.ConfigurationImp
}

type demoImp struct {
}

func (this *demoImp) HelloNoReponse(_ context.Context, req *rpc.HelloRequest) error {
	fmt.Println("receive hello request3", req)
	return nil
}

func main() {
	server.Run(&Config{}, func(microCtx server.MicroContext) {
		rpc.RegisterK8sdemoForHTTP(&demoImp{})
	})
}
