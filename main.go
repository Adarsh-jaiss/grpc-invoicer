package main

import (
	"log"
	"net"
	"context"

	"google.golang.org/grpc"
	"github.com/adarsh-jaiss/grpc-invoicer/invoicer"
)

type myInvoicerServer struct{
	invoicer.UnimplementedInvoicerServer

}

func(s myInvoicerServer) Create(ctx context.Context, in *invoicer.CreateRequest, opts ...grpc.CallOption) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		pdf: []byte("test"),
		Docx : []byte("test"),
	}, nil
}

func main()  {
	lis,err := net.Listen("tcp","3000")
	if err!= nil{
		log.Fatal(err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myInvoicerServer{

	}

	invoicer.RegisterInvoicerServer(serverRegistrar,service)
	err = serverRegistrar.Serve(lis)
	if err!= nil{
		log.Fatal(err)
	}
	
}