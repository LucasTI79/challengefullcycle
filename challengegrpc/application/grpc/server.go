package grpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/jinzhu/gorm"
	"github.com/lucasti79/desafiogrpc/application/grpc/pb"
	"github.com/lucasti79/desafiogrpc/application/usecase"
	"github.com/lucasti79/desafiogrpc/infra/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	productRepository := repository.ProductRepositoryDb{Db: database}
	productUseCase := usecase.ProductUseCase{ProductRepository: productRepository}
	productGrpcService := NewProductGrpcService(productUseCase)
	pb.RegisterProductServiceServer(grpcServer, productGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

	go signalsListener(grpcServer)

	log.Printf("gRPC server has been started on port %d", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}

func signalsListener(server *grpc.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	_ = <-sigs

	log.Printf("Gracefully stopping server...")
	server.GracefulStop()
}
