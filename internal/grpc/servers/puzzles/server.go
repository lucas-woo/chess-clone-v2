package puzzlegrpc

import (
	puzzlesv1 "github.com/lucas-woo/chess-clone-v2/api/puzzles/v1"
)

type Server struct {
	puzzlesv1.UnimplementedPuzzlesServiceServer
}

func NewServer () *Server {
	return &Server{

	};
}