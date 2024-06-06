package main

import (
	"context"
	"fmt"
	cataloguev1 "github.com/NursultanBayzakov/protos/gen/go/catalogue"
	ssov1 "github.com/NursultanBayzakov/protos/gen/go/sso"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := ssov1.RegisterAuthHandlerFromEndpoint(ctx, mux, "0.0.0.0:44044", opts)
	if err != nil {
		return
	}

	err = cataloguev1.RegisterCatalogueServiceHandlerFromEndpoint(ctx, mux, "0.0.0.0:44045", opts)
	if err != nil {
		return
	}

	log.Println("HTTP server listening on :8888")
	err = http.ListenAndServe(":8888", mux)
	if err != nil {
		log.Fatalf(fmt.Sprintf("failed to listen %v", err))
		return
	}
}
