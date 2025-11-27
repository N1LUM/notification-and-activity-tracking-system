package configs

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(addr string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           addr,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		TLSConfig:      initTLSConfig(),
	}

	appCertPath := getEnvOrFatal("GATEWAY_CERT_PATH")
	appCertKeyPath := getEnvOrFatal("GATEWAY_KEY_PATH")

	return s.httpServer.ListenAndServeTLS(appCertPath, appCertKeyPath)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func GetAppAddress() string {
	return fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
}

func initTLSConfig() *tls.Config {
	caCertPath := getEnvOrFatal("CA_CERT_PATH")

	caCert, err := os.ReadFile(caCertPath)
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	appCertPath := getEnvOrFatal("GATEWAY_CERT_PATH")
	appCertKeyPath := getEnvOrFatal("GATEWAY_KEY_PATH")

	cert, err := tls.LoadX509KeyPair(appCertPath, appCertKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.VerifyClientCertIfGiven,
	}
}

func getEnvOrFatal(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("%s is not set", key)
	}
	return val
}
