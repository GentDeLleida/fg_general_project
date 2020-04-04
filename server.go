package main
import ( pb "game" 
		"fmt" 
		"net"
		"google.golang.org/grpc"
		"flag"
	)

func main (){
	port    := flag.Int("port", 50051, "The server port")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFightingGameServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}