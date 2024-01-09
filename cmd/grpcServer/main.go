package main

import (
	"database/sql"
	"net"

	"github.com/kameikay/grpc-example/internal/database"
	"github.com/kameikay/grpc-example/internal/pb"
	"github.com/kameikay/grpc-example/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	listener, err := net.Listen("tcp", ":50051")
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
