package server

import (
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

func (d *datasourceServer) AddDatasources(in *ds.AddDatasourceRequest, stream ds.Datasource_AddDatasourcesServer) error {
	fmt.Println("Add Datasource Request Data --- ", in)
	if err := stream.Send(&ds.MessageResponse{Message: "Success"}); err != nil {
		return err
	}
	return nil
}

func (d *datasourceServer) ListDatasources(in *ds.ListDatasourceRequest, stream ds.Datasource_ListDatasourcesServer) error {
	fmt.Println("List Datasource Requested Data --- ", in)
	//To Store Datasource
	var datasource []*ds.ListAllDatasource
	//Hardcoded Data
	dbData := []*ds.ListAllDatasource{
		{DsId: "1", DsDatadomain: "org1.com", DsName: "Google", DsDescription: "Google Data Domain", DsType: "MYSQL", DsVersion: "8", DsKey: "hjdhgdfgb3645634545", CreatedAt: "10-07-2022 12:15:34"},
		{DsId: "2", DsDatadomain: "org2.com", DsName: "AWS", DsDescription: "Amazon Data Domain", DsType: "MYSQL", DsVersion: "8", DsKey: "hjdtyrydfgb3645634545", CreatedAt: "10-07-2022 12:15:34"},
		{DsId: "3", DsDatadomain: "org3.com", DsName: "Google", DsDescription: "Google Data Domain", DsType: "MYSQL", DsVersion: "5", DsKey: "hjdhgdfgb3645634545", CreatedAt: "10-07-2022 12:15:34"},
		{DsId: "4", DsDatadomain: "org4.com", DsName: "Google", DsDescription: "Google Data Domain", DsType: "MYSQL", DsVersion: "8", DsKey: "hjdhgdfgb3645634545", CreatedAt: "10-07-2022 12:15:34"},
		{DsId: "5", DsDatadomain: "org5.com", DsName: "AWS", DsDescription: "Amazon Data Domain", DsType: "MYSQL", DsVersion: "8", DsKey: "hjfghfgdfgb3645634545", CreatedAt: "10-07-2022 12:15:34"},
		{DsId: "6", DsDatadomain: "org6.com", DsName: "AWS", DsDescription: "Amazon Data Domain", DsType: "MYSQL", DsVersion: "8", DsKey: "hjdhgdfgb36456909gbjgh45", CreatedAt: "10-07-2022 12:15:34"},
		{DsId: "7", DsDatadomain: "org7.com", DsName: "Google", DsDescription: "Google Data Domain", DsType: "MYSQL", DsVersion: "8", DsKey: "hjdhgdfgbhjghjghj4545", CreatedAt: "10-07-2022 12:15:34"},
		{DsId: "8", DsDatadomain: "org8.com", DsName: "AWS", DsDescription: "Amazon Data Domain", DsType: "MYSQL", DsVersion: "5", DsKey: "ghjghjghjghjghj78678678", CreatedAt: "10-07-2022 12:15:34"},
		{DsId: "9", DsDatadomain: "org9.com", DsName: "Google", DsDescription: "Google Data Domain", DsType: "MYSQL", DsVersion: "8", DsKey: "hjdhgdfgb3645634545", CreatedAt: "10-07-2022 12:15:34"},
		{DsId: "10", DsDatadomain: "org10.com", DsName: "AWS", DsDescription: "Amazon Data Domain", DsType: "MYSQL", DsVersion: "8", DsKey: "ghjghjghjgh4456fghfgh", CreatedAt: "10-07-2022 12:15:34"},
		{DsId: "11", DsDatadomain: "org11.com", DsName: "AWS", DsDescription: "Amazon Data Domain", DsType: "MYSQL", DsVersion: "8", DsKey: "fghfghfgh4556fghdfhh", CreatedAt: "10-07-2022 12:15:34"},
		{DsId: "12", DsDatadomain: "org12.com", DsName: "Google", DsDescription: "Google Data Domain", DsType: "MYSQL", DsVersion: "8", DsKey: "fghfgy5668ghh5654", CreatedAt: "10-07-2022 12:15:34"},
	}
	//Send Response to Client
	datasource = append(datasource, dbData...)
	if err := stream.Send(&ds.ListDatasourceResponse{ListAllDatasources: datasource}); err != nil {
		return err
	}
	return nil
}

func (d *datasourceServer) DeleteDatasources(in *ds.DeleteDatasourceRequest, stream ds.Datasource_DeleteDatasourcesServer) error {
	fmt.Println("Delete Datasource Requested Data --- ", in)
	if err := stream.Send(&ds.MessageResponse{Message: "Success"}); err != nil {
		return err
	}
	return nil
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
