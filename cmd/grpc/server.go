package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"messaging-interface/config"
	"messaging-interface/domain/notifications"
	pb "messaging-interface/domain/notifications"
	"messaging-interface/pkg/dbconnect"

	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) SendEmail(ctx context.Context, req *pb.EmailRequest) (*pb.EmailResponse, error) {

	fmt.Println("Received Email Request:", req)
	return &pb.EmailResponse{
		Success: true,
		Message: "Email sent successfully",
	}, nil

}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}
	config.AppConfig = cfg

	dbg := dbconnect.DBConfig{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		Dbname:   cfg.Postgres.DbName,
		Sslmode:  cfg.Postgres.SSLMode,
	}
	db, err := dbconnect.ConnectSqlx(dbg)
	if err != nil {
		fmt.Println(err)
	}

	if db == nil {
		panic("db not connected")
	}

	srv := grpc.NewServer()
	notifications.RegisterRouteGRPC(srv, db)

	log.Println("Register RouteGRPC ...")

	listen, err := net.Listen("tcp", cfg.App.GrpcPort)
	if err != nil {
		panic(err)
	}

	log.Println("running GRPC server at port", cfg.App.GrpcPort)
	err = srv.Serve(listen)
	if err != nil {
		panic(err)
	}

}
