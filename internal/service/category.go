package service

import (
	"context"
	"github.com/bianavic/fullcycle_gRPC/internal/database"
	"github.com/bianavic/fullcycle_gRPC/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(db database.Category) *CategoryService {
	return &CategoryService{CategoryDB: db}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description) // in vem da request
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating category: %v", err)
	}
	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: *category.Description,
	}
	return categoryResponse, err
}
