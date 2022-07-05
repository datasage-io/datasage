package server

import (
	"sync"

	dpclassifcation "github.com/datasage-io/datasage/src/classifiers"
	logger "github.com/datasage-io/datasage/src/logger"

	pb "github.com/datasage-io/datasage/src/proto/datasource"
	"github.com/google/uuid"
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

type DatasourceServer struct {
	pb.UnimplementedDatasourceServer
}

/*
func (s *DatasourceServer) AddDatasources(ctx context.Context, in *pb.AddDatasourceRequest, opts ...grpc.CallOption) (pb.Datasource_AddDatasourcesClient, error) {
	log.Trace().Msg("AddDatasources")
	return nil, nil
}
func (s *DatasourceServer) ListDatasources(ctx context.Context, in *pb.ListDatasourceRequest, opts ...grpc.CallOption) (pb.Datasource_ListDatasourcesClient, error) {
	return nil, nil
}
func (s *DatasourceServer) DeleteDatasources(ctx context.Context, in *pb.DeleteDatasourceRequest, opts ...grpc.CallOption) (pb.Datasource_DeleteDatasourcesClient, error) {
	return nil, nil
}
*/

func (s *DatasourceServer) AddDatasources(in *pb.AddDatasourceRequest, ser pb.Datasource_AddDatasourcesServer) error {
	log.Trace().Msg(in.String())

	dpDataSource := dpclassifcation.DpDataSource{
		Datadomain:   in.GetDataDomain(),
		Dsname:       in.GetDsName(),
		Dsdecription: in.GetDsDescription(),
		Dstype:       in.GetDsType(),
		Dsversion:    in.GetDsVersion(),
		Host:         in.GetHost(),
		Port:         in.GetPort(),
		User:         in.GetUser(),
		Password:     in.GetPassword(),
		DsKey:        uuid.New().String(),
	}

	go dpclassifcation.Run(dpDataSource)
	return nil
}

func (s *DatasourceServer) ListDatasources(*pb.ListDatasourceRequest, pb.Datasource_ListDatasourcesServer) error {
	log.Trace().Msg("ListDatasources")

	return nil
}
func (s *DatasourceServer) DeleteDatasources(*pb.DeleteDatasourceRequest, pb.Datasource_DeleteDatasourcesServer) error {
	log.Trace().Msg("DeleteDatasources")
	return nil

}

// ================= //
// == gRPC Server == //
// ================= //

//GetNewServer - gRPC Server
func GetNewServer() *grpc.Server {
	log.Info().Msg("gRPC Server Started....")
	s := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	//Create Server Instance
	//Register gRPC Server
	pb.RegisterDatasourceServer(s, &DatasourceServer{})

	reflection.Register(s)

	return s
}
