package main

import (
  "fmt"
  "log"
  "net"
  "net/http"
  "strings"

  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "golang.org/x/net/context"

  "github.com/ldb358/GRPCMTGInventory/api"
  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials"
  "google.golang.org/grpc/metadata"
)

// private type for Context keys
type contextKey int
const (
  clientIDKey contextKey = iota
)

func credMatcher(headerToken string) (mdName string, ok bool) {
    if headerToken == "Token" {
        return headerToken, true
    }
    return "", false
}

// authenticateAgent check the client credentials
func authenticateClient(ctx context.Context, s *api.Server) (string, error) {
    if md, ok := metadata.FromIncomingContext(ctx); ok {
        clientToken := strings.Join(md["token"], "")
        if clientToken != "098f6bcd4621d373cade4e832627b4f6" {
        return "", fmt.Errorf("unknown user %s", clientToken)
        }
        log.Printf("authenticated client: %s", clientToken)
        return "42", nil
    }
    return "", fmt.Errorf("missing credentials")
}
// unaryInterceptor calls authenticateClient with current context
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
    s, ok := info.Server.(*api.Server)
    if !ok {
        return nil, fmt.Errorf("unable to cast server")
    }
    clientID, err := authenticateClient(ctx, s)
    if err != nil {
        return nil, err
    }
    ctx = context.WithValue(ctx, clientIDKey, clientID)
    return handler(ctx, req)
}

func startGRPCServer(address, certFile, keyFile string) error {
    lis, err := net.Listen("tcp", address)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    // create a server instance
    s := api.Server{}

    // Create the TLS credentials
    creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
    if err != nil {
        log.Fatalf("Could not load TLS keys: %s", err)
    }
    // Create an array of gRPC options with the credentials
    opts := []grpc.ServerOption{grpc.Creds(creds),
        grpc.UnaryInterceptor(unaryInterceptor)}

    // create a gRPC server object
    grpcServer := grpc.NewServer(opts...)
    // attach the Ping service to the server
    api.RegisterMTGCardServiceServer(grpcServer, &s)
    // start the server
    log.Printf("Starting HTTP/2 gRPC server on %s", address)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %s", err)
    }
    return nil
}

func startRESTServer(address, grpcAddress, certFile string) error {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()
    mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(credMatcher))
    creds, err := credentials.NewClientTLSFromFile(certFile, "")
    if err != nil {
        return fmt.Errorf("could not load TLS certificate: %s", err)
    }
    // Setup the client gRPC options
    opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
    // Register ping
    err = api.RegisterMTGCardServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
    if err != nil {
        return fmt.Errorf("could not register service Ping: %s", err)
    }
    log.Printf("starting HTTP/1.1 REST server on %s", address)
    http.ListenAndServe(address, mux)
    return nil
}

// main start a gRPC server and waits for connection
func main() {
    // create a listener on TCP port 7777
    grpcAddress := fmt.Sprintf("%s:%d", "localhost", 7777)
    restAddress := fmt.Sprintf("%s:%d", "localhost", 7778)
    certFile := "cert/server.crt"
    keyFile := "cert/server.key"
    // fire the gRPC server in a goroutine
    go func() {
      err := startGRPCServer(grpcAddress, certFile, keyFile)
      if err != nil {
        log.Fatalf("failed to start gRPC server: %s", err)
      }
    }()
    // fire the REST server in a goroutine
    go func() {
      err := startRESTServer(restAddress, grpcAddress, certFile)
      if err != nil {
        log.Fatalf("failed to start gRPC server: %s", err)
      }
    }()
    // infinite loop
    log.Printf("Entering infinite loop")
    select {}
}