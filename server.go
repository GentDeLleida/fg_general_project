package main
import ( pb "game" 
		"fmt" 
		"net"
		"google.golang.org/grpc"
		"flag"
		"context"
	)

func (TypeShit) StartGame(context.Context, *pb.StartRequest) (*pb.StartResponse, error){
	return nil, nil
}

func (TypeShit) Game(pb.FightingGame_GameServer) error{
	return nil
}

type TypeShit struct{}	

func main (){
	port    := flag.Int("port", 50051, "The server port")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	srv := TypeShit{}
	pb.RegisterFightingGameServer(grpcServer, srv)
	fmt.Println("Startig Server")
	grpcServer.Serve(lis)
}