package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)



func main(){
	var port = ":7777"

	listener, err := net.Listen("tcp", port)
	if err != nil {
			log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}

	grpcService := grpc.NewServer()


	log.Printf("Servidor gRPC rodando na porta:%v", port)
	if err := grpcService.Serve(listener); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}