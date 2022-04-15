package rpc_server

import (
	"dtstack.com/dtstack/easymatrix/addons/easyfiler/pkg/proto"
	"dtstack.com/dtstack/easymatrix/go-common/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	Port    string
	Root    string
	WithDB  bool
	Rate    int
	Isziped bool
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		log.Errorf("failed to listen, err: %v", err)
		return err
	}
	srv := grpc.NewServer()
	proto.RegisterTransferServiceServer(srv, &FileTransferService{Root: s.Root, WithDB: s.WithDB, Rate: s.Rate, IsZiped: s.Isziped})
	reflection.Register(srv)
	if err := srv.Serve(listener); err != nil {
		log.Errorf("failed to serve, err: %v", err)
		return err
	}
	return nil
}
