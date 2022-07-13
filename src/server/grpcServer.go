package server

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	dpclassifcation "github.com/datasage-io/datasage/src/classifiers"
	logger "github.com/datasage-io/datasage/src/logger"

	ds "github.com/datasage-io/datasage/src/proto/datasource"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

const PortNumber = "8089"

var log *zerolog.Logger = logger.GetInstance()
var wg sync.WaitGroup

// ======================= //
// == Datasource Service == //
// ===================== //

type DatasourceServer struct {
	ds.UnimplementedDatasourceServer
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
/*
func (s *DatasourceServer) AddDatasources(in *pb.AddDatasourceRequest, stream pb.Datasource_AddDatasourcesServer) error {
	log.Trace().Msg("AddDatasources Enter")

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

	dpclassifcation.Run(dpDataSource)

	log.Trace().Msg("AddDatasources Exit")
	return nil
}

func (s *DatasourceServer) ListDatasources(*pb.ListDatasourceRequest, pb.Datasource_ListDatasourcesServer) error {
	log.Trace().Msg("ListDatasources")
	dpclassifcation.ListDatasources()
	return nil
}
func (s *DatasourceServer) DeleteDatasources(*pb.DeleteDatasourceRequest, pb.Datasource_DeleteDatasourcesServer) error {
	log.Trace().Msg("DeleteDatasources")
	dpclassifcation.DeleteDatasources()
	return nil

}
*/

func (d *DatasourceServer) AddDatasource(ctx context.Context, in *ds.AddRequest) (*ds.MessageResponse, error) {
	fmt.Println("Add Datasource Request --- ", in)
	dpDataSource := dpclassifcation.DpDataSource{
		Datadomain:   in.GetDataDomain(),
		Dsname:       in.GetName(),
		Dsdecription: in.GetDescription(),
		Dstype:       in.GetType(),
		Dsversion:    in.GetVersion(),
		Host:         in.GetHost(),
		Port:         in.GetPort(),
		User:         in.GetUser(),
		Password:     in.GetPassword(),
		DsKey:        uuid.New().String(),
	}
	go dpclassifcation.Run(dpDataSource)
	fmt.Println("Add Datasource Request Data sending response ", in)
	return &ds.MessageResponse{Message: "Success"}, nil
}
func (d *DatasourceServer) ListDatasource(ctx context.Context, in *ds.ListRequest) (*ds.ListResponse, error) {
	fmt.Println("List Datasource Request ", in)
	datasources, err := dpclassifcation.ListDatasources()
	if err != nil {
		fmt.Println("ListDatasources error  ")
	}
	datasourcesOut := []*ds.ListAll{}
	for _, datasource := range datasources {
		fmt.Println(datasource)

		outDS := &ds.ListAll{
			Id:          strconv.Itoa(datasource.ID),
			Datadomain:  datasource.Datadomain,
			Name:        datasource.Dsname,
			Description: datasource.Dsdecription,
			Type:        datasource.Dstype,
			Version:     datasource.Dsversion,
			Key:         datasource.DsKey,
			CreatedAt:   datasource.CreatedAt,
		}
		datasourcesOut = append(datasourcesOut, outDS)
	}
	return &ds.ListResponse{ListAllDatasources: datasourcesOut, Count: int64(len(datasourcesOut))}, nil
}
func (d *DatasourceServer) DeleteDatasource(ctx context.Context, in *ds.DeleteRequest) (*ds.MessageResponse, error) {
	fmt.Println("Delete Datasource Request --- ", in)
	var ids []int64
	arrayIds := in.GetId()
	for i := range arrayIds {
		element := arrayIds[i]
		id, err := strconv.ParseInt(element, 10, 64)
		if err != nil {
			return &ds.MessageResponse{Message: "incorrect input"}, nil
		}
		ids = append(ids, id)
	}
	statusDelete := dpclassifcation.DeleteDatasource(ids)
	if statusDelete == true {
		return &ds.MessageResponse{Message: "Delete sucessful"}, nil
	}
	return &ds.MessageResponse{Message: "Delete failed"}, nil
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
	ds.RegisterDatasourceServer(s, &DatasourceServer{})
	return s
}
