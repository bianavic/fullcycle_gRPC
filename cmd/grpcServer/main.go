package main

import (
	"database/sql"
	"net"

	"github.com/bianavic/fullcycle_gRPC/internal/database"
	"github.com/bianavic/fullcycle_gRPC/internal/pb"
	"github.com/bianavic/fullcycle_gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := initializeDatabase(db); err != nil {
		panic(err)
	}

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)
	// le e processa a propria informacao

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer) // usando reflection para usar evans

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}

func initializeDatabase(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS categories (
		id TEXT PRIMARY KEY,
		name TEXT,
		description TEXT
	);`

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	return err
}
