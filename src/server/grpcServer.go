package server

import (
	"sync"

	logger "github.com/datasage-io/datasage/src/logger"
	ds "github.com/datasage-io/datasage/src/proto/datasource"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

const PortNumber = "8089"

var log *zerolog.Logger = logger.GetInstance()
var wg sync.WaitGroup

// ======================= //
// == Datasource Service == //
// ===================== //

type datasourceServer struct {
	ds.DatasourceServer
}

// ================= //
// == gRPC Server == //
// ================= //

//GetNewServer - gRPC Server
func GetNewServer() *grpc.Server {

	s := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	reflection.Register(s)

	return s
}
