package grpc

import (
	"context"

	"github.com/brianvoe/gofakeit"
	"github.com/tweeker88/chat-server-oleg/pkg/chat_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	chat_v1.UnimplementedChatV1Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Create(_ context.Context, _ *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	return &chat_v1.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *Server) SendMessage(_ context.Context, _ *chat_v1.SendMessageRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *Server) Delete(_ context.Context, _ *chat_v1.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
