package main

import (
	"context"
	"google.golang.org/grpc"
	"lab_gRPC_speed_test/request"
	"log"
	"time"
)

const (
	conNumber    = 16
	clientNumber = 10000
)

func main() {
	var conn [conNumber]*grpc.ClientConn
	var err error
	for i := 0; i < conNumber; i++ {
		conn[i], err = grpc.Dial(":9000", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
	}
	//conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("did not connect: %s", err)
	//}
	//defer conn.Close()

	//c := request.NewSimpleServerClient(conn)
	//response, err := c.Request(context.Background(), &request.Empty{Body: ""})
	//if err != nil {
	//	log.Fatalf("Error when calling IsThere: %s", err)
	//}
	//fmt.Printf("%s\n", response.Body)
	for i := 0; i < conNumber; i++ {
		go func(conn *grpc.ClientConn) {
			for k := 0; k < clientNumber; k++ {
				c := request.NewSimpleServerClient(conn)
				go func(c *request.SimpleServerClient) {
					for {
						(*c).Request(context.Background(), &request.Empty{Body: ""})
					}
				}(&c)
			}
		}(conn[i])
	}
	time.Sleep(5 * time.Minute)
}
