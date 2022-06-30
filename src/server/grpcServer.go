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

func (d *datasourceServer) AddDatasources(in *ds.AddDatasourceRequest, stream ds.Datasource_AddDatasourcesServer) error {
	return nil
}

func (d *datasourceServer) ListDatasources(in *ds.ListDatasourceRequest, stream ds.Datasource_ListDatasourcesServer) error {
	return nil
}

func (d *datasourceServer) DeleteDatasources(in *ds.DeleteDatasourceRequest, stream ds.Datasource_DeleteDatasourcesServer) error {
	return nil
}

// ================= //
// == gRPC Server == //
// ================= //

//GetNewServer - gRPC Server
func GetNewServer() *grpc.Server {
	log.Info().Msg("// ================= //")
	log.Info().Msg("// == gRPC Server Started == //")
	log.Info().Msg("// ================= //")
	s := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	//Create Server Instance
	datasourceServer := &datasourceServer{}

	//Register gRPC Server
	ds.RegisterDatasourceServer(s, datasourceServer)

	reflection.Register(s)

	return s
}
