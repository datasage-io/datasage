package main

import (
	integrations "github.com/datasage-io/datasage/src/integrations"
	grpcServer "github.com/datasage-io/datasage/src/server"
)

func main() {

	//start integration component dependent servers
	go integrations.RunServer()
	//Run a gRPC Server for CLI command processing
	grpcServer.RunServer()

}
