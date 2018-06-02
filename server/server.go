package server

import (
	"log"
	"net"

	"os"

	"github.com/morganwu277/kvdb/db"
	"github.com/morganwu277/kvdb/server/pb"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement the server defined in protobuf 
type server struct {
	db *db.KVDB
}

// implements Read(context.Context, *KVRequest) (*KVResponse, error)
func (s *server) Read(ctx context.Context, in *pb.KVRequest) (*pb.KVResponse, error) {
	value, err := s.db.Read(in.Key)
	if err != nil {
		log.Printf("error reading key: %v, error: %v \n ", in.Key, err)
		return &pb.KVResponse{Key: in.Key, Value: "", ErrMsg: err.Error()}, err
	}
	return &pb.KVResponse{Key: in.Key, Value: value, ErrMsg: ""}, nil
}

// implements Write(context.Context, *KVRequest) (*KVResponse, error)
func (s *server) Write(ctx context.Context, in *pb.KVRequest) (*pb.KVResponse, error) {
	err := s.db.Write(in.Key, in.Value)
	if err != nil {
		log.Printf("error writing key: %v, value: %v, error: %v \n ", in.Key, in.Value, err)
		return &pb.KVResponse{Key: in.Key, Value: in.Value, ErrMsg: err.Error()}, err
	}
	return &pb.KVResponse{Key: in.Key, Value: in.Value, ErrMsg: ""}, nil
}

func StartServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	dbPath := "/tmp/lev-test"
	db, err := db.Init(dbPath)
	if err != nil {
		log.Fatalf("failed to init DB from path: %v \n ", dbPath)
		os.Exit(1)
	}
	ks := &server{
		db: db,
	}
	pb.RegisterKVDBServiceServer(s, ks)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
