package server

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"sync"

	logger "github.com/datasage-io/datasage/src/logger"
	"github.com/datasage-io/datasage/src/utils/constants"
	"github.com/golang/protobuf/ptypes/empty"

	classpb "github.com/datasage-io/datasage/src/proto/class"
	ds "github.com/datasage-io/datasage/src/proto/datasource"
	tagpb "github.com/datasage-io/datasage/src/proto/tag"
	"github.com/datasage-io/datasage/src/storage"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

const PortNumber = "8089"

var log *zerolog.Logger = logger.GetInstance()
var wg sync.WaitGroup

// ======================================== //
// == RPC Server ( Datasource Class Tag) == //
// ======================================= //

type DatasourceServer struct {
	ds.UnimplementedDatasourceServer
}

type TagServer struct {
	tagpb.UnimplementedTagServer
}

type ClassServer struct {
	classpb.UnimplementedClassServer
}

// ====================//
// == Class Service == //
// =================== //

func (d *ClassServer) AddClass(ctx context.Context, in *classpb.CreateRequest) (*classpb.MessageResponse, error) {
	fmt.Println("AddClass : ", in)
	st, err := storage.GetStorageInstance()
	if err != nil {
		log.Error().Err(err).Msg("Internal Error")
	} else {
		rules := in.GetTag()
		log.Debug().Msgf("rules %v", rules)
		for _, rule := range rules {
			log.Debug().Msgf("Rule  className %v", rule)

			err := st.AddClass(in.GetDescription(), rule, in.GetName())
			if err != nil {
				log.Error().Err(err).Msg("Internal Error")
			}
		}
		return &classpb.MessageResponse{Message: "Class Added sucessfully"}, nil
	}
	return &classpb.MessageResponse{Message: "Error in adding Class "}, nil

}
func (d *ClassServer) ListClass(ctx context.Context, in *classpb.ListRequest) (*classpb.ListResponse, error) {
	fmt.Println("ListClass : ", in)

	st, err := storage.GetStorageInstance()
	if err != nil {
		log.Error().Err(err).Msg("Internal Error")
	}

	classes, _ := st.GetClasses()
	classesOut := []*classpb.ClassResponse{}
	for _, class := range classes {
		log.Debug().Msgf("ListTag %v", class)
		classOut := &classpb.ClassResponse{
			Id:          strconv.Itoa(class.Id),
			Name:        class.Rule,
			Description: class.Description,
			Tag:         "",
			CreatedAt:   "",
		}
		classesOut = append(classesOut, classOut)

	}
	return &classpb.ListResponse{ClassResponse: classesOut, Count: int64(len(classesOut))}, nil

	return nil, nil
}
func (d *ClassServer) DeleteClass(ctx context.Context, in *classpb.DeleteRequest) (*classpb.MessageResponse, error) {
	fmt.Println("DeleteClass : ", in)
	st, err := storage.GetStorageInstance()
	if err != nil {
		log.Error().Err(err).Msg("Internal Error")
	}
	var ids []int64
	arrayIds := in.GetId()
	for i := range arrayIds {
		element := arrayIds[i]
		id, err := strconv.ParseInt(element, 10, 64)
		if err != nil {
			return &classpb.MessageResponse{Message: "incorrect input"}, nil
		}
		ids = append(ids, id)
	}
	statusDelete, err := st.DeleteClasses(ids)
	if err != nil {
		log.Error().Err(err).Msg("Internal Error")
	}
	if statusDelete == true {
		return &classpb.MessageResponse{Message: "Delete sucessful"}, nil
	}
	return &classpb.MessageResponse{Message: "Delete failed"}, nil

}

// ====================//
// == Tag   Service == //
// =================== //
func (d *TagServer) AddTag(ctx context.Context, in *tagpb.AddRequest) (*tagpb.MessageResponse, error) {
	log.Debug().Msgf("AddTag %v", in)

	st, err := storage.GetStorageInstance()
	if err != nil {
		log.Error().Err(err).Msg("Internal Error")
	} else {
		var classesAsscociated []string
		classes := in.GetClass()

		log.Debug().Msgf("classes %v", classes)

		for _, className := range classes {
			log.Debug().Msgf("AddTag  className %v", className)
			classesAsscociated = append(classesAsscociated, className)
		}
		log.Debug().Msgf("array %v", classesAsscociated)

		err1 := st.AddTag(in.GetName(), in.GetDescription(), classesAsscociated)
		if err1 != nil {
			log.Error().Err(err).Msg("Internal Error")
			return &tagpb.MessageResponse{Message: "Error in adding Tag"}, nil
		}
	}
	return &tagpb.MessageResponse{Message: "Tag Added sucessfully"}, nil

}
func (d *TagServer) ListTag(ctx context.Context, in *tagpb.ListRequest) (*tagpb.ListResponse, error) {
	log.Debug().Msgf("ListTag %v", in)

	st, err := storage.GetStorageInstance()
	if err != nil {
		log.Error().Err(err).Msg("Internal Error")
	}

	tags, _ := st.GetTags()
	tagsOut := []*tagpb.TagResponse{}

	for _, tag := range tags {
		log.Debug().Msgf("ListTag %v", tag)
		classes := []string{tag.Rule}
		tagOut := &tagpb.TagResponse{
			Id:          strconv.Itoa(tag.Id),
			Name:        tag.TagName,
			Description: tag.Description,
			Class:       classes,
			CreatedAt:   "",
		}
		tagsOut = append(tagsOut, tagOut)

	}
	return &tagpb.ListResponse{TagResponse: tagsOut, Count: int64(len(tagsOut))}, nil
}
func (d *TagServer) DeleteTag(ctx context.Context, in *tagpb.DeleteRequest) (*tagpb.MessageResponse, error) {
	log.Debug().Msgf("DeleteTag %v", in)
	st, err := storage.GetStorageInstance()
	if err != nil {
		log.Error().Err(err).Msg("Internal Error")
	}
	var ids []int64
	arrayIds := in.GetId()
	for i := range arrayIds {
		element := arrayIds[i]
		id, err := strconv.ParseInt(element, 10, 64)
		if err != nil {
			return &tagpb.MessageResponse{Message: "incorrect input"}, nil
		}
		ids = append(ids, id)
	}
	statusDelete, err := st.DeleteTags(ids)
	if err != nil {
		log.Error().Err(err).Msg("Internal Error")
	}
	if statusDelete == true {
		return &tagpb.MessageResponse{Message: "Delete sucessful"}, nil
	}
	return &tagpb.MessageResponse{Message: "Delete failed"}, nil

}



// ====================//
// == Datasource Service == //
// =================== //

func (d *DatasourceServer) AddDatasource(ctx context.Context, in *ds.AddRequest) (*ds.AddResponse, error) {
	fmt.Println("Add Datasource Request --- ", in)
	st, err := storage.GetStorageInstance()
	if err != nil {
		log.Error().Err(err).Msg("Internal Error")
	}

	storageDpDataSourceObj := storage.DpDataSource{
		ID:           -1,
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

	err = st.AddDataSource(storageDpDataSourceObj)
	if err != nil {
		return &ds.AddResponse{StatusCode: codes.Internal.String(), Message: ""}, status.Error(codes.Internal, "Internal Error")
	}
	return &ds.AddResponse{StatusCode: codes.OK.String(), Message: "Data Source added for Scaning"}, nil
}
func (d *DatasourceServer) ListDatasource(ctx context.Context, in *ds.ListRequest) (*ds.ListResponse, error) {
	fmt.Println("List Datasource Request ", in)
	st, err := storage.GetStorageInstance()
	if err != nil {
		log.Error().Err(err).Msg("Internal Error")
	}

	datasources, err := st.GetDataSources()
	if err != nil {
		fmt.Println("Datasources not found ")
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
func (d *DatasourceServer) DeleteDatasource(ctx context.Context, in *ds.DeleteRequest) (*ds.DeleteResponse, error) {
	fmt.Println("Delete Datasource Request --- ", in)
	st, err := storage.GetStorageInstance()
	if err != nil {
		log.Error().Err(err).Msg("Internal Error")
	}

	var ids []int64
	arrayIds := in.GetId()
	for i := range arrayIds {
		element := arrayIds[i]
		id, err := strconv.ParseInt(element, 10, 64)
		if err != nil {
			return &ds.DeleteResponse{StatusCode: codes.InvalidArgument.String(), Message: "incorrect input"}, status.Error(codes.InvalidArgument, "incorrect input")
		}
		ids = append(ids, id)
	}

	statusDelete, err := st.DeleteDataSources(ids)
	if err != nil {
		log.Error().Err(err).Msg("Internal Error")
	}
	if statusDelete {
		return &ds.DeleteResponse{StatusCode: codes.OK.String(), Message: "Delete sucessful"}, nil
	}
	return &ds.DeleteResponse{StatusCode: codes.Unknown.String(), Message: "Delete failed"}, status.Error(codes.Unknown, "Delete failed")
}

func (d *DatasourceServer) LogDatasource(ctx context.Context, in *ds.LogRequest) (*ds.LogResponse, error) {
	fmt.Println("Request for log -- ", in)
	return nil, nil
}

func (d *DatasourceServer) Scan(ctx context.Context, in *ds.ScanRequest) (*ds.ScanResponse, error) {
	fmt.Println("Request for Scan - ", in)
	return &ds.ScanResponse{StatusCode: codes.OK.String(), Message: "Scan Completed"}, nil
}

func (d *DatasourceServer) GetStatus(ctx context.Context, in *ds.StatusRequest) (*ds.StatusResponse, error) {
	fmt.Println("Request for Status - Datasource Name - ", in)
	return &ds.StatusResponse{StatusCode: codes.OK.String(), DsStatus: constants.DataSourceInitialScanCompleted}, nil
}

func (d *DatasourceServer) GetRecommendedPolicy(ctx context.Context, in *empty.Empty) (*ds.RecommendedpolicyResponse, error) {
	fmt.Println("GetRecommendedPolicy - ", in)
	//RecommendedPolicy
	var recommedPolicy = []string{"GDPR Audit", "PII Audit", "SOC 2 Audit", "HIPAA Audit", "UDI Audit"}
	return &ds.RecommendedpolicyResponse{StatusCode: codes.OK.String(), PolicyName: recommedPolicy}, nil
}

func (d *DatasourceServer) ApplyPolicy(ctx context.Context, in *ds.ApplyPolicyRequest) (*ds.ApplyPolicyResponse, error) {
	fmt.Println("Apply Recommended Policy Request - ", in)
	return &ds.ApplyPolicyResponse{StatusCode: codes.OK.String(), Message: "Recommended Policy Applied"}, nil
}

// ================= //
// == gRPC Server == //
// ================= //

func RunServer() {

	listen, err := net.Listen("tcp", ":"+PortNumber)
	if err != nil {
		log.Error().Msgf("gRPC server failed to listen : %v", err)
	}

	s := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	//Create Server Instance
	//Register gRPC Server
	ds.RegisterDatasourceServer(s, &DatasourceServer{})
	tagpb.RegisterTagServer(s, &TagServer{})
	classpb.RegisterClassServer(s, &ClassServer{})

	//Start service
	log.Info().Msgf("gRPC server on %s port started", PortNumber)
	if err := s.Serve(listen); err != nil {
		log.Error().Msgf("Failed to serve: %v", err)
	}
}
