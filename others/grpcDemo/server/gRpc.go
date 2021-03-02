package server

import (
	"goProjects/daylearning/others/grpcDemo/order"
	"google.golang.org/grpc"
	"net"
)

// GrpcServer 为订单服务实现 gRPC 服务
type GrpcServer struct {
	server   *grpc.Server
	errCh    chan error
	listener net.Listener
}

//NewGrpcServer 是一个创建 GrpcServer 的便捷函数
func NewGrpcServer(service order.OrderServiceServer, port string) (GrpcServer, error) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return GrpcServer{}, err
	}
	server := grpc.NewServer()
	order.RegisterOrderServiceServer(server, service)
	return GrpcServer{
		server:   server,
		listener: lis,
		errCh:    make(chan error),
	}, nil
}

// Start 在后台启动服务，将任何错误传入错误通道
func (g GrpcServer) Start() {
	go func() {
		g.errCh <- g.server.Serve(g.listener)
	}()
}

// Stop 停止 gRPC 服务
func (g GrpcServer) Stop() {
	g.server.GracefulStop()
}

//Error 返回服务的错误通道
func (g GrpcServer) Error() chan error {
	return g.errCh
}
