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
	return nil
}

func (d *datasourceServer) ListDatasources(in *ds.ListDatasourceRequest, stream ds.Datasource_ListDatasourcesServer) error {
	fmt.Println("List Datasource Requested Data --- ", in)
	return nil
}

func (d *datasourceServer) DeleteDatasources(in *ds.DeleteDatasourceRequest, stream ds.Datasource_DeleteDatasourcesServer) error {
	fmt.Println("Delete Datasource Requested Data --- ", in)
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
	return nil
}

func (t *tagServer) ListTag(in *tag.ListTagRequest, stream tag.Tag_ListTagServer) error {
	fmt.Println("List Tag Requested Data --- ", in)
	return nil
}

func (t *tagServer) DeleteTag(in *tag.DeleteTagRequest, stream tag.Tag_DeleteTagServer) error {
	fmt.Println("Delete Tag Requested Data --- ", in)
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
	return nil
}

func (c *classServer) ListClass(in *class.ListClassRequest, stream class.Class_ListClassServer) error {
	fmt.Println("List Class Requested Data --- ", in)
	return nil
}

func (c *classServer) DeleteClass(in *class.DeleteClassRequest, stream class.Class_DeleteClassServer) error {
	fmt.Println("Delete CLass Requested Data --- ", in)
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
