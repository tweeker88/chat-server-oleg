package builder

import (
	grpc2 "github.com/tweeker88/chat-server-oleg/internal/transport/grpc"
	"github.com/tweeker88/chat-server-oleg/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func BuildServer() *grpc.Server {
	srv := grpc.NewServer()
	reflection.Register(srv)

	my := grpc2.NewServer()

	chat_v1.RegisterChatV1Server(srv, my)

	return srv
}
