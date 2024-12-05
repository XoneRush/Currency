package main

import (
	"net"

	protos "github.com/XoneRush/Currency/protos"
	"github.com/XoneRush/Currency/server"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrency(log)
	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Error("Unable to listen", "error", err)
	}

	gs.Serve(l)
}
