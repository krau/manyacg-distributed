package client

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"github.com/krau/manyacg/core/proto"
	"github.com/krau/manyacg/storage/config"
	"github.com/krau/manyacg/storage/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var ArtworkClient proto.ArtworkServiceClient

func init() {
	// 公钥中读取和解析公钥/私钥对
	pair, err := tls.LoadX509KeyPair(config.Cfg.App.CertFile, config.Cfg.App.KeyFile)
	if err != nil {
		logger.L.Fatalf("Failed to load certificates: %s", err)
		return
	}
	// 创建一组根证书
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(config.Cfg.App.CaFile)
	if err != nil {
		logger.L.Fatalf("Failed to load certificates: %s", err)
		return
	}
	// 解析证书
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		logger.L.Fatalf("Failed to load certificates")
		return
	}
	cred := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{pair},
		ServerName:   config.Cfg.App.ServerName,
		RootCAs:      certPool,
	})
	conn, err := grpc.Dial(config.Cfg.App.GrpcAddr, grpc.WithTransportCredentials(cred))
	if err != nil {
		logger.L.Fatalf("Failed to dial: %s", err)
		return
	}
	ArtworkClient = proto.NewArtworkServiceClient(conn)
}
