package server

import (
	"context"
	"fmt"
	"sync"

	logger "github.com/datasage-io/datasage/src/logger"
	class "github.com/datasage-io/datasage/src/proto/class"
	ds "github.com/datasage-io/datasage/src/proto/datasource"
	tag "github.com/datasage-io/datasage/src/proto/tag"
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

func (d *datasourceServer) AddDatasource(ctx context.Context, in *ds.AddRequest) (*ds.MessageResponse, error) {
	fmt.Println("Add Datasource Request --- ", in)
	return &ds.MessageResponse{Message: "Success"}, nil
}
func (d *datasourceServer) ListDatasource(ctx context.Context, in *ds.ListRequest) (*ds.ListResponse, error) {
	fmt.Println("List Datasource Request --- ", in)
	var datasource []*ds.ListAll
	//Hardcoded Data
	dbData := []*ds.ListAll{
		{Id: "1", Datadomain: "org1.com", Name: "Google", Description: "Google Data Domain", Type: "MYSQL", Version: "8", Key: "hjdhgdfgb3645634545", CreatedAt: "10-07-2022 12:15:34"},
		{Id: "2", Datadomain: "org2.com", Name: "AWS", Description: "Amazon Data Domain", Type: "MYSQL", Version: "8", Key: "hjdtyrydfgb3645634545", CreatedAt: "10-07-2022 12:15:34"},
		{Id: "3", Datadomain: "org3.com", Name: "Google", Description: "Google Data Domain", Type: "MYSQL", Version: "5", Key: "hjdhgdfgb3645634545", CreatedAt: "10-07-2022 12:15:34"},
		{Id: "4", Datadomain: "org4.com", Name: "Google", Description: "Google Data Domain", Type: "MYSQL", Version: "8", Key: "hjdhgdfgb3645634545", CreatedAt: "10-07-2022 12:15:34"},
		{Id: "5", Datadomain: "org5.com", Name: "AWS", Description: "Amazon Data Domain", Type: "MYSQL", Version: "8", Key: "hjfghfgdfgb3645634545", CreatedAt: "10-07-2022 12:15:34"},
		{Id: "6", Datadomain: "org6.com", Name: "AWS", Description: "Amazon Data Domain", Type: "MYSQL", Version: "8", Key: "hjdhgdfgb36456909gbjgh45", CreatedAt: "10-07-2022 12:15:34"},
		{Id: "7", Datadomain: "org7.com", Name: "Google", Description: "Google Data Domain", Type: "MYSQL", Version: "8", Key: "hjdhgdfgbhjghjghj4545", CreatedAt: "10-07-2022 12:15:34"},
		{Id: "8", Datadomain: "org8.com", Name: "AWS", Description: "Amazon Data Domain", Type: "MYSQL", Version: "5", Key: "ghjghjghjghjghj78678678", CreatedAt: "10-07-2022 12:15:34"},
		{Id: "9", Datadomain: "org9.com", Name: "Google", Description: "Google Data Domain", Type: "MYSQL", Version: "8", Key: "hjdhgdfgb3645634545", CreatedAt: "10-07-2022 12:15:34"},
		{Id: "10", Datadomain: "org10.com", Name: "AWS", Description: "Amazon Data Domain", Type: "MYSQL", Version: "8", Key: "ghjghjghjgh4456fghfgh", CreatedAt: "10-07-2022 12:15:34"},
		{Id: "11", Datadomain: "org11.com", Name: "AWS", Description: "Amazon Data Domain", Type: "MYSQL", Version: "8", Key: "fghfghfgh4556fghdfhh", CreatedAt: "10-07-2022 12:15:34"},
		{Id: "12", Datadomain: "org12.com", Name: "Google", Description: "Google Data Domain", Type: "MYSQL", Version: "8", Key: "fghfgy5668ghh5654", CreatedAt: "10-07-2022 12:15:34"},
	}
	//Send Response to Client
	datasource = append(datasource, dbData...)
	return &ds.ListResponse{ListAllDatasources: datasource, Count: 10}, nil

}
func (d *datasourceServer) DeleteDatasource(ctx context.Context, in *ds.DeleteRequest) (*ds.MessageResponse, error) {
	fmt.Println("Delete Datasource Request --- ", in)
	return &ds.MessageResponse{Message: "Success"}, nil
}

// ======================= //
// == Tag Service == //
// ===================== //

type tagServer struct {
	tag.TagServer
}

func (t *tagServer) AddTag(ctx context.Context, in *tag.AddRequest) (*tag.MessageResponse, error) {
	fmt.Println("Delete Tag Request --- ", in)
	return &tag.MessageResponse{Message: "Success"}, nil
}
func (t *tagServer) ListTag(ctx context.Context, in *tag.ListRequest) (*tag.ListResponse, error) {
	fmt.Println("List Tag Request --- ", in)
	//Hardcoded Data
	dbData := []*tag.TagResponse{
		{Id: "1", Name: "PII-3", Description: "Personal Identofiable Information", Class: []string{"Postal Address,Bank Account"}, CreatedAt: "10-07-2022 12:15:34"},
		{Id: "2", Name: "PII", Description: "Personal Identofiable Information", Class: []string{"Bank Account"}, CreatedAt: "10-07-2022 12:15:34"},
		{Id: "3", Name: "GDPR", Description: "General Data Protection Regulation", Class: []string{"Credit Card"}, CreatedAt: "10-07-2022 12:15:34"},
		{Id: "4", Name: "HIPAA", Description: "Portability Insurance And Accountablity Act Payment Card", Class: []string{"Health Card"}, CreatedAt: "10-07-2022 12:15:34"},
		{Id: "5", Name: "PCI-DSS", Description: "Industry Data Security Standard", Class: []string{"Bank Account"}, CreatedAt: "10-07-2022 12:15:34"},
		{Id: "6", Name: "PHI", Description: "Protected Health Information", Class: []string{"Drug Enforcement Agency Registration Number"}, CreatedAt: "10-07-2022 12:15:34"},
	}
	//Send Response to Client
	return &tag.ListResponse{TagResponse: dbData, Count: 6}, nil
}
func (t *tagServer) DeleteTag(ctx context.Context, in *tag.DeleteRequest) (*tag.MessageResponse, error) {
	fmt.Println("Delete Tag Request --- ", in)
	return &tag.MessageResponse{Message: "Success"}, nil
}

// ======================= //
// == Class Service == //
// ===================== //

type classServer struct {
	class.ClassServer
}

func (c *classServer) AddClass(ctx context.Context, in *class.CreateRequest) (*class.MessageResponse, error) {
	fmt.Println("Delete Class Request --- ", in)
	return &class.MessageResponse{Message: "Success"}, nil
}
func (c *classServer) ListClass(ctx context.Context, in *class.ListRequest) (*class.ListResponse, error) {
	fmt.Println("List Class Request --- ", in)
	//Hardcoded Data
	dbData := []*class.ClassResponse{
		{Id: "1", Name: "Indian Moblie Number", Description: "Indian Moblie Number", Tag: []string{"PII-2"}, CreatedAt: "10-07-2022 12:15:34"},
		{Id: "2", Name: "Passport Number", Description: "Contains Passport Number", Tag: []string{"PII"}, CreatedAt: "10-07-2022 12:15:34"},
		{Id: "3", Name: "Social Security Number", Description: "Contains Social Security Number", Tag: []string{"PII"}, CreatedAt: "10-07-2022 12:15:34"},
		{Id: "4", Name: "Drivers License Number", Description: "Contains Drivers License ID Number", Tag: []string{"PII"}, CreatedAt: "10-07-2022 12:15:34"},
		{Id: "5", Name: "Phone Number", Description: "Contains Phone Number", Tag: []string{"PII"}, CreatedAt: "10-07-2022 12:15:34"},
		{Id: "6", Name: "AWS secrets", Description: "Contains AWS Secrets", Tag: []string{"GDPR"}, CreatedAt: "10-07-2022 12:15:34"},
	}
	//Send Response to Client
	return &class.ListResponse{ClassResponse: dbData, Count: 6}, nil

}
func (c *classServer) DeleteClass(ctx context.Context, in *class.DeleteRequest) (*class.MessageResponse, error) {
	fmt.Println("Delete Class Request --- ", in)
	return &class.MessageResponse{Message: "Success"}, nil
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
	tagServer := &tagServer{}
	classServer := &classServer{}

	//Register gRPC Server
	ds.RegisterDatasourceServer(s, datasourceServer)
	tag.RegisterTagServer(s, tagServer)
	class.RegisterClassServer(s, classServer)

	reflection.Register(s)

	return s
}
