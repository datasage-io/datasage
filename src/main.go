package main

import (
	"net"

	grpcServer "github.com/datasage-io/datasage/src/server"
	log "github.com/rs/zerolog/log"
)

func main() {

	//Run a gRPC Server for CLI command processing
	listen, err := net.Listen("tcp", ":"+grpcServer.PortNumber)
	if err != nil {
		log.Error().Msgf("gRPC server failed to listen : %v", err)
	}
	server := grpcServer.GetNewServer()
	//Start service
	log.Info().Msgf("gRPC server on %s port started", grpcServer.PortNumber)
	if err := server.Serve(listen); err != nil {
		log.Error().Msgf("Failed to serve: %v", err)
	}

}
