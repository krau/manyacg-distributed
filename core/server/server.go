package server

import (
	"crypto/tls"
	"crypto/x509"
	"net"
	"os"

	"github.com/krau/Picture-collector/core/config"
	"github.com/krau/Picture-collector/core/logger"
	"github.com/krau/Picture-collector/core/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ArtworkServer struct {
	proto.UnimplementedArtworkServiceServer
}

var server = &ArtworkServer{}

func StartGrpcServer() {
	lis, err := net.Listen("tcp", config.Cfg.App.Address)
	if err != nil {
		logger.L.Fatalf("Failed to listen: %s", err)
		return
	}
	pair, err := tls.LoadX509KeyPair(config.Cfg.App.CertFile, config.Cfg.App.KeyFile)
	if err != nil {
		logger.L.Fatalf("Failed to load certificates: %s", err)
		return
	}

	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(config.Cfg.App.CaFile)
	if err != nil {
		logger.L.Fatalf("Failed to load certificates: %s", err)
		return
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		logger.L.Fatalf("Failed to load certificates: %s", err)
		return
	}
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{pair},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	s := grpc.NewServer(grpc.Creds(creds))
	proto.RegisterArtworkServiceServer(s, server)
	logger.L.Noticef("Grpc server listen on %s with TLS", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.L.Fatalf("Failed to serve: %s", err)
		return
	}
}
