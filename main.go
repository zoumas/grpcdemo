package main

import (
	"context"
	"github.com/zoumas/grpcdemo/invoicer"
	"google.golang.org/grpc"
	"log"
	"net"
)

type InvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s *InvoicerServer) Create(ctx context.Context, r *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte(r.From),
		Docx: []byte(r.To),
	}, nil
}

func main() {
	listener := unwrap(net.Listen("tcp", ":8080"))

	serviceRegistrar := grpc.NewServer()
	invoicerServer := &InvoicerServer{}
	invoicer.RegisterInvoicerServer(serviceRegistrar, invoicerServer)
	log.Fatal(serviceRegistrar.Serve(listener))
}

func unwrap[T any](v T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return v
}
