package rpc

import "context"

type HelloRequest struct {
	Send int `bson:"send" json:"send"`
}
type HelloResponse struct {
	Ok bool `bson:"ok" json:"ok"`
}
type K8sdemo interface {
	HelloNoReponse(context.Context, *HelloRequest) error
}
