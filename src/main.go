package main

import (
	"net"

	"github.com/datasage-io/datasage/src/classifiers"
	grpcserver "github.com/datasage-io/datasage/src/server"
	"github.com/rs/zerolog/log"
)

func main() {

	classifiers.Run()

	//Create Server
	listen, err := net.Listen("tcp", ":"+grpcserver.PortNumber)
	if err != nil {
		log.Error().Msgf("gRPC server failed to listen : %v", err)
	}

	server := grpcserver.GetNewServer()
	//Start service
	log.Info().Msgf("gRPC server on %s port started", grpcserver.PortNumber)
	if err := server.Serve(listen); err != nil {
		log.Error().Msgf("Failed to serve: %v", err)
	}

}
