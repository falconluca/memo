package memo

import (
	"context"
	"github.com/shaohsiung/memo/api/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	memo := &Item{
		Title:       req.Item.Title,
		Description: req.Item.Description,
		RemindAt:    req.Item.RemindAt.AsTime(),
	}
	resp := &protobuf.CreateResponse{}
	if err := s.r.Create(memo); err != nil {
		return resp, status.Errorf(codes.Internal, "create memo failed: %v", err)
	}
	resp.Id = memo.ID
	return resp, nil
}

func (s Server) Update(ctx context.Context, req *protobuf.UpdateRequest) (*protobuf.UpdateResponse, error) {
	memoId := req.Item.Id
	resp := &protobuf.UpdateResponse{}
	if _, err := s.r.Get(memoId); err != nil {
		return resp, status.Error(codes.NotFound, "memo not found")
	}

	memo := &Item{
		ID:          memoId,
		Title:       req.Item.Title,
		Description: req.Item.Description,
		RemindAt:    req.Item.RemindAt.AsTime(),
	}
	if err := s.r.Update(memo); err != nil {
		return resp, status.Errorf(codes.Internal, "update memo failed: %v", err)
	}
	resp.Updated = true
	return resp, nil
}

func (s Server) Delete(ctx context.Context, req *protobuf.DeleteRequest) (*protobuf.DeleteResponse, error) {
	memoId := req.Id
	resp := &protobuf.DeleteResponse{}
	if _, err := s.r.Get(memoId); err != nil {
		return resp, status.Errorf(codes.NotFound, "memo not found")
	}

	if err := s.r.Delete(memoId); err != nil {
		return resp, status.Errorf(codes.Internal, "delete memo failed: %v", err)
	}
	resp.Deleted = true
	return resp, nil
}

func (s Server) Get(ctx context.Context, req *protobuf.GetRequest) (*protobuf.GetResponse, error) {
	resp := &protobuf.GetResponse{}
	memo, err := s.r.Get(req.Id)
	if err != nil {
		return resp, status.Errorf(codes.NotFound, "memo not found")
	}

	resp.Item = &protobuf.Item{
		Id:          memo.ID,
		Title:       memo.Title,
		Description: memo.Description,
		RemindAt:    timestamppb.New(memo.RemindAt),
	}
	return resp, nil
}

func (s Server) List(ctx context.Context, req *protobuf.ListRequest) (*protobuf.ListResponse, error) {
	resp := &protobuf.ListResponse{}
	memos, err := s.r.List()
	if err != nil {
		return resp, status.Errorf(codes.Internal, "list memos failed: %v", err)
	}
	for _, memo := range memos {
		resp.Items = append(resp.Items, &protobuf.Item{
			Id:          memo.ID,
			Title:       memo.Title,
			Description: memo.Description,
			RemindAt:    timestamppb.New(memo.RemindAt),
		})
	}
	return resp, nil
}
