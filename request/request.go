package request

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
)

type Server struct {
	OkStatusCounter prometheus.Counter
}

func (s *Server) mustEmbedUnimplementedSimpleServerServer() {

}

func (s *Server) Request(ctx context.Context, empty *Empty) (*Empty, error) {
	s.OkStatusCounter.Inc()
	return &Empty{Body: "hi"}, nil
}
