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

func (t *tagServer) AddTag(in *tag.CreateTagRequest, stream tag.Tag_AddTagServer) error {
	fmt.Println("Add Tag Request Data --- ", in)
	if err := stream.Send(&tag.TagMessageResponse{Message: "Success"}); err != nil {
		return err
	}
	return nil
}

func (t *tagServer) ListTag(in *tag.ListTagRequest, stream tag.Tag_ListTagServer) error {
	fmt.Println("List Tag Requested Data --- ", in)
	//To Store Tag
	var tagdata []*tag.TagResponse
	//Hardcoded Data
	dbData := []*tag.TagResponse{
		{TagId: "1", TagName: "PII-3", TagDescription: "Personal Identofiable Information", TagClass: "Postal Address", CreatedAt: "10-07-2022 12:15:34"},
		{TagId: "2", TagName: "PII", TagDescription: "Personal Identofiable Information", TagClass: "Bank Account", CreatedAt: "10-07-2022 12:15:34"},
		{TagId: "3", TagName: "GDPR", TagDescription: "General Data Protection Regulation", TagClass: "Credit Card", CreatedAt: "10-07-2022 12:15:34"},
		{TagId: "4", TagName: "HIPAA", TagDescription: "Portability Insurance And Accountablity Act Payment Card", TagClass: "Health Card", CreatedAt: "10-07-2022 12:15:34"},
		{TagId: "5", TagName: "PCI-DSS", TagDescription: "Industry Data Security Standard", TagClass: "Bank Account", CreatedAt: "10-07-2022 12:15:34"},
		{TagId: "6", TagName: "PHI", TagDescription: "Protected Health Information", TagClass: "Drug Enforcement Agency Registration Number", CreatedAt: "10-07-2022 12:15:34"},
	}
	//Send Response to Client
	tagdata = append(tagdata, dbData...)
	if err := stream.Send(&tag.ListTagResponse{TagResponse: tagdata}); err != nil {
		return err
	}
	return nil
}

func (t *tagServer) DeleteTag(in *tag.DeleteTagRequest, stream tag.Tag_DeleteTagServer) error {
	fmt.Println("Delete Tag Requested Data --- ", in)
	if err := stream.Send(&tag.TagMessageResponse{Message: "Success"}); err != nil {
		return err
	}
	return nil
}

// ======================= //
// == Class Service == //
// ===================== //

type classServer struct {
	class.ClassServer
}

func (c *classServer) AddClass(in *class.CreateClassRequest, stream class.Class_AddClassServer) error {
	fmt.Println("Add Class Request Data --- ", in)
	if err := stream.Send(&class.ClassMessageResponse{Message: "Success"}); err != nil {
		return err
	}
	return nil
}

func (c *classServer) ListClass(in *class.ListClassRequest, stream class.Class_ListClassServer) error {
	fmt.Println("List Class Requested Data --- ", in)
	//To Store Class
	var classdata []*class.ClassResponse
	//Hardcoded Data
	dbData := []*class.ClassResponse{
		{ClassId: "1", ClassName: "Indian Moblie Number", ClassDescription: "Indian Moblie Number", ClassTag: "PII-2", CreatedAt: "10-07-2022 12:15:34"},
		{ClassId: "1", ClassName: "Passport Number", ClassDescription: "Contains Passport Number", ClassTag: "PII", CreatedAt: "10-07-2022 12:15:34"},
		{ClassId: "1", ClassName: "Social Security Number", ClassDescription: "Contains Social Security Number", ClassTag: "PII", CreatedAt: "10-07-2022 12:15:34"},
		{ClassId: "1", ClassName: "Drivers License Number", ClassDescription: "Contains Drivers License ID Number", ClassTag: "PII", CreatedAt: "10-07-2022 12:15:34"},
		{ClassId: "1", ClassName: "Phone Number", ClassDescription: "Contains Phone Number", ClassTag: "PII", CreatedAt: "10-07-2022 12:15:34"},
		{ClassId: "1", ClassName: "AWS secrets", ClassDescription: "Contains AWS Secrets", ClassTag: "GDPR", CreatedAt: "10-07-2022 12:15:34"},
	}
	//Send Response to Client
	classdata = append(classdata, dbData...)
	if err := stream.Send(&class.ListClassResponse{ClassResponse: classdata}); err != nil {
		return err
	}
	return nil
}

func (c *classServer) DeleteClass(in *class.DeleteClassRequest, stream class.Class_DeleteClassServer) error {
	fmt.Println("Delete CLass Requested Data --- ", in)
	if err := stream.Send(&class.ClassMessageResponse{Message: "Success"}); err != nil {
		return err
	}
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
	tagServer := &tagServer{}
	classServer := &classServer{}

	//Register gRPC Server
	ds.RegisterDatasourceServer(s, datasourceServer)
	tag.RegisterTagServer(s, tagServer)
	class.RegisterClassServer(s, classServer)

	reflection.Register(s)

	return s
}
