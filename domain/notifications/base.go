package notifications

import (
	"github.com/jmoiron/sqlx"
	grpc "google.golang.org/grpc"
)

func RegisterRouteGRPC(server *grpc.Server, db *sqlx.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewHandler(svc)

	RegisterMessagingServiceServer(server, handler)
	// RegisterAuthServer(server, handler)
}
