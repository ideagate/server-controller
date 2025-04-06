package grpc

import (
	"fmt"
	"net"
	"net/http"

	"github.com/ideagate/core/config"
	"github.com/ideagate/model/gen-go/dashboard"
	apprepositorysql "github.com/ideagate/server-controller/domain/application/repository/sql"
	appusecase "github.com/ideagate/server-controller/domain/application/usecase"
	projectrepositorysql "github.com/ideagate/server-controller/domain/project/repository/sql"
	projectusecase "github.com/ideagate/server-controller/domain/project/usecase"
	"github.com/ideagate/server-controller/infrastructure"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Action(_ *cli.Context) error {
	// Load configuration
	if err := config.Load("."); err != nil {
		return fmt.Errorf("failed to load configuration: %v", err)
	}

	// Initialize TCP connection
	lisGrpc, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	lisGrpcWeb, err := net.Listen("tcp", ":50052")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// Initialize infrastructure
	infra, err := infrastructure.NewInfrastructure(config.Get())
	if err != nil {
		return fmt.Errorf("failed to initialize infrastructure: %v", err)
	}

	// Register gRPC services
	s := grpc.NewServer()
	dashboard.RegisterDashboardServiceServer(s, NewDashboardServiceServer(infra))

	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		fmt.Println("gRPC server is running on port 50051")
		if err := s.Serve(lisGrpc); err != nil {
			fmt.Errorf("failed to serve: %v", err)
		}
	}()

	wrappedGrpcWeb := grpcweb.WrapServer(s)
	httpServer := &http.Server{
		Handler: middleware(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			if wrappedGrpcWeb.IsGrpcWebRequest(req) {
				wrappedGrpcWeb.ServeHTTP(resp, req)
				return
			}
			// Fall back to other servers.
			http.DefaultServeMux.ServeHTTP(resp, req)
		})),
	}

	go func() {
		fmt.Println("gRPC web server is running on port 50052")
		if err := httpServer.Serve(lisGrpcWeb); err != nil {
			fmt.Errorf("failed to serve: %v", err)
		}
	}()

	select {}
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Disposition, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-grpc-web")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Max-Age", "600")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

type DashboardServiceServer struct {
	dashboard.UnimplementedDashboardServiceServer

	usecaseProject     projectusecase.ProjectUsecase
	usecaseApplication appusecase.ApplicationUsecase
}

func NewDashboardServiceServer(infra *infrastructure.Infrastructure) *DashboardServiceServer {
	// Initialize repository
	repoProject := projectrepositorysql.NewProjectRepository(infra.Postgres)
	repoApplication := apprepositorysql.NewApplicationRepository(infra.Postgres)

	// Initialize usecase
	usecaseProject := projectusecase.NewProjectUsecase(repoProject)
	usecaseApplication := appusecase.NewApplicationUsecase(repoApplication)

	return &DashboardServiceServer{
		usecaseProject:     usecaseProject,
		usecaseApplication: usecaseApplication,
	}
}
