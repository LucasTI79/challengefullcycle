package grpc

import (
	"context"

	"github.com/lucasti79/desafiogrpc/application/grpc/pb"
	"github.com/lucasti79/desafiogrpc/application/usecase"
)

type ProductGrpcService struct {
	ProductUseCase usecase.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func (p *ProductGrpcService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product, err := p.ProductUseCase.RegisterProduct(in.Name, in.Description, in.Price)
	if err != nil {
		return &pb.CreateProductResponse{
			Product: &pb.Product{},
		}, err
	}

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		},
	}, nil
}

func (p *ProductGrpcService) FindProducts(ctx context.Context, in *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	products, err := p.ProductUseCase.FindProducts()

	if err != nil {
		return &pb.FindProductsResponse{
			Products: []*pb.Product{},
		}, err
	}

	response := &pb.FindProductsResponse{}
	for _, product := range products {
		response.Products = append(response.Products, &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		})
	}

	return response, nil
}

func NewProductGrpcService(usecase usecase.ProductUseCase) *ProductGrpcService {
	return &ProductGrpcService{
		ProductUseCase: usecase,
	}
}
