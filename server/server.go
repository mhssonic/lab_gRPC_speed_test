package main

import (
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"lab_gRPC_speed_test/request"
	"log"
	"net"
	"net/http"
)

var okStatusCounter prometheus.Counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ok_request_count",
		Help: "Number of 200",
	},
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := request.Server{OkStatusCounter: okStatusCounter}

	grpcServer := grpc.NewServer()
	request.RegisterSimpleServerServer(grpcServer, &s)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	prometheus.MustRegister(okStatusCounter)
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", 3333),
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error running http server: %s\n", err)
		}
	}
}
