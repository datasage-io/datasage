package integrations

//go:generate echo "Generating gRPC files"
//go:generate protoc --proto_path=grpc_config grpc_config/datasage.proto --go-grpc_out=. --go_out=.

import (
	"context"
	"log"
	"time"

	"github.com/datasage-io/datasage/src/integrations/grpc_config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StreamLogToGRPC(Log string, grpcConfigs []GRPCLogConfig) error {
	for _, config := range grpcConfigs {
		var conn *grpc.ClientConn
		log.Printf("[GRPC] Dialing %s:%s \n", config.Host, config.Port)
		conn, err := grpc.Dial(config.Host+":"+config.Port,
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %s", err)
			return err
		}
		defer conn.Close()
		c := grpc_config.NewDataSageServerClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		if _, err = c.LogSend(ctx, &grpc_config.Log{Body: Log}); err != nil {
			log.Printf("could not send the data: %v", err)
			return err
		}
		log.Printf("[GRPC] Log send to %s:%s\n", config.Host, config.Port)
	}
	return nil
}
