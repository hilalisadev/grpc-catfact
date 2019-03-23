package catfact

import (
	"golang.org/x/net/context"
	"log"
)

// Server contains all available procedures
type Server struct {
}

// Get receiver to return a random cat fact
func (s *Server) Get(ctx context.Context) (*CatFactResponse, error) {
	log.Printf("get fact")
	return &CatFactResponse{Fact: "cats are pretty"}, nil
}
