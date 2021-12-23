package memo

import (
	"context"
	"github.com/shaohsiung/memo/api/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	protobuf.UnimplementedMemoServer
	r Repository
}

func NewServer(r Repository) protobuf.MemoServer {
	return &Server{
		r: r,
	}
}

func (s Server) Create(ctx context.Context, req *protobuf.CreateRequest) (*protobuf.CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func (s Server) Update(ctx context.Context, req *protobuf.UpdateRequest) (*protobuf.UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func (s Server) Delete(ctx context.Context, req *protobuf.DeleteRequest) (*protobuf.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func (s Server) Get(ctx context.Context, req *protobuf.GetRequest) (*protobuf.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func (s Server) List(ctx context.Context, req *protobuf.ListRequest) (*protobuf.ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
