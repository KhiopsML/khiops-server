package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	khiops "gitlab.tech.orange/khiops-tools/docker-khiops/go/gen/proto"
	selfcert "gitlab.tech.orange/khiops-tools/docker-khiops/go/insecure"
	"gitlab.tech.orange/khiops-tools/docker-khiops/go/server"
	"google.golang.org/grpc/credentials/insecure"

	// Static files
	_ "gitlab.tech.orange/khiops-tools/docker-khiops/go/statik"
)

// For the record: size without openapi
// 15315908 uncompressed
//  2867656 compressed

// getOpenAPIHandler serves an OpenAPI UI.
// Adapted from https://github.com/philips/grpc-gateway-example/blob/a269bcb5931ca92be0ceae6130ac27ae89582ecc/cmd/serve.go#L63
func getOpenAPIHandler() http.Handler {
	err := mime.AddExtensionType(".svg", "image/svg+xml")
	if err != nil {
		panic("failed registering mimetype for svg: " + err.Error())
	}

	statikFS, err := fs.New()
	if err != nil {
		panic("creating OpenAPI filesystem: " + err.Error())
	}

	return http.FileServer(statikFS)
}

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}

// Main function
func Main() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)
	grpcport := os.Getenv("GRPC_PORT")
	if grpcport == "" {
		grpcport = "10000"
	}
	addr := "0.0.0.0:" + grpcport
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	plaintext := strings.ToLower(os.Getenv("GRPC_PLAINTEXT")) == "true"
	s := grpc.NewServer()
	if !plaintext {
		s = grpc.NewServer(
			// TODO: Replace with your own certificate!
			grpc.Creds(credentials.NewServerTLSFromCert(&selfcert.Cert)),
		)
	}
	b := server.New(ctx)
	khiops.RegisterKhiopsServer(s, b)
	khiops.RegisterOperationsServer(s, b)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	// Serve gRPC Server
	if plaintext {
		log.Info("Serving gRPC on http://", addr)
	} else {
		log.Info("Serving gRPC on https://", addr)
	}

	go func() {
		log.Fatal(s.Serve(lis))
	}()

	// See https://github.com/grpc/grpc/blob/master/doc/naming.md
	// for gRPC naming standard information.
	dialAddr := fmt.Sprintf("dns:///localhost:%s", grpcport)
	// Create a client connection to the gRPC Server we just started.
	// This is where the gRPC-Gateway proxies the requests.
	scheme := grpc.WithTransportCredentials(insecure.NewCredentials())
	if !plaintext {
		scheme = grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(selfcert.CertPool, ""))
	}
	conn, err := grpc.DialContext(
		context.Background(),
		dialAddr,
		scheme,
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux(
	//runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonpb),
	)
	err = khiops.RegisterKhiopsHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway for khiops service:", err)
	}
	err = khiops.RegisterOperationsHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway for operations service:", err)
	}

	oa := getOpenAPIHandler()

	// Empty parameters mean use the TLS Config specified with the server.
	servehttp := strings.ToLower(os.Getenv("SERVE_HTTP")) == "true"
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "11000"
	}
	gatewayAddr := "0.0.0.0:" + port
	gwServer := &http.Server{
		Addr: gatewayAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/v1") {
				gwmux.ServeHTTP(w, r)
				return
			}
			oa.ServeHTTP(w, r)
		}),
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: time.Second,
	}

	go func() {
		<-c
		log.Warning("SIGINT caught, exiting.")
		gwServer.Close()
		s.Stop()
		cancel()
	}()

	if servehttp {
		log.Info("Serving REST API on http://", gatewayAddr)
		log.Fatalln(gwServer.ListenAndServe())
	} else {
		gwServer.TLSConfig = &tls.Config{
			MinVersion:   tls.VersionTLS12,
			Certificates: []tls.Certificate{selfcert.Cert},
		}
		log.Info("Serving REST API on https://", gatewayAddr)
		log.Fatalln(gwServer.ListenAndServeTLS("", ""))
	}

	return nil
}
